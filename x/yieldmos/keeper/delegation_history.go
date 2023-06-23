package keeper

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/yieldmos/types"
	"time"
)

// SetDelegationHistory set a specific delegationHistory in the store from its index
func (k Keeper) SetDelegationHistory(ctx sdk.Context, delegationHistory types.DelegationHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationHistoryKeyPrefix))
	b := k.cdc.MustMarshal(&delegationHistory)
	store.Set(types.DelegationHistoryKey(
		delegationHistory.Address,
	), b)
}

// GetDelegationHistory returns a delegationHistory from its index
func (k Keeper) GetDelegationHistory(
	ctx sdk.Context,
	address string,

) (val types.DelegationHistory, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationHistoryKeyPrefix))

	b := store.Get(types.DelegationHistoryKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegationHistory removes a delegationHistory from the store
func (k Keeper) RemoveDelegationHistory(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationHistoryKeyPrefix))
	store.Delete(types.DelegationHistoryKey(
		address,
	))
}

// GetAllDelegationHistory returns all delegationHistory
func (k Keeper) GetAllDelegationHistory(ctx sdk.Context) (list []types.DelegationHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DelegationHistoryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegationHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CheckDelegationHistoryRecords checks if the accounts DelegationHistory record needs to be updated
func (k Keeper) CheckDelegationHistoryRecords(ctx sdk.Context, delAddr sdk.AccAddress) error {
	if k.stakingKeeper == nil {
		panic("stakingKeeper is nil")
	}

	var (
		delegationHistory types.DelegationHistory
		found             bool
		err               error
	)

	//current staked amount
	delegatedAmount := k.CalcTotalDelegatedAmount(ctx, delAddr)

	//if no DelegationHistory exists, create a new one
	delegationHistory, found = k.GetDelegationHistory(ctx, delAddr.String())
	if !found {
		delegationHistory = k.NewDelegationHistory(ctx, delAddr, delegatedAmount)
	}

	//calculates the difference between the sum of DelegationTimestamps balances in the DelegationHistory and the current amount staked
	difference := k.CalcDelegationHistoryDifference(delegatedAmount, delegationHistory)

	//on a positive difference only add a DelegationTimestamp to a DelegationHistory
	if difference.IsPositive() {
		delegationHistory = k.AddDelegationTimestamp(ctx, difference, delegationHistory)
	}

	//on a negative difference it might adjust and/or remove a DelegationTimestamp in a DelegationHistory
	if difference.IsNegative() {
		delegationHistory, err = k.AdjustDelegationTimestamps(delegationHistory, difference)
		if err != nil {
			return err
		}
	}

	//adjustDelegationTimestamps might set some DelegationTimestamps to 0, if so remove it from the DelegationHistory
	prunedDelegationHistory := k.PruneDelegationHistory(delegationHistory)

	//save changes made to the DelegationHistory
	k.SetDelegationHistory(ctx, prunedDelegationHistory)

	return nil
}

// CalcTotalDelegatedAmount gets all delegations shares, sums them and converts to bond denom
func (k Keeper) CalcTotalDelegatedAmount(ctx sdk.Context, delAddr sdk.AccAddress) sdk.Int {
	delegatedAmount := sdk.NewDec(0)
	delegations := k.stakingKeeper.GetAllDelegatorDelegations(ctx, delAddr)

	for _, delegation := range delegations {
		val, found := k.stakingKeeper.GetValidator(ctx, delegation.GetValidatorAddr())
		if !found {
			panic("unable to find validator: " + delegation.GetValidatorAddr().String())
		}

		delegatedAmount = delegatedAmount.Add(val.TokensFromShares(delegation.Shares))
	}

	return delegatedAmount.TruncateInt()
}

// CalcDelegationHistoryDifference calculates the difference between the sum of bonded amounts in a DelegationHistory and the delegation amount
func (k Keeper) CalcDelegationHistoryDifference(delegationAmount sdk.Int, delegationHistory types.DelegationHistory) sdk.Int {
	delegatedAmountHistory := sdk.NewInt(0)
	for _, delegationHistory := range delegationHistory.GetHistory() {
		delegatedAmountHistory = delegatedAmountHistory.Add(delegationHistory.GetBalance().Amount)
	}

	return delegationAmount.Sub(delegatedAmountHistory)
}

// AddDelegationTimestamp adds a DelegationTimestamp to an existing DelegationHistory record
func (k Keeper) AddDelegationTimestamp(ctx sdk.Context, amount sdk.Int, delegationHistory types.DelegationHistory) types.DelegationHistory {
	delegationTimestamp := k.NewDelegationTimestamp(ctx, amount)

	delegationHistory.History = append(delegationHistory.History, &delegationTimestamp)

	return delegationHistory
}

// AdjustDelegationTimestamps removes DelegationsTimestamp(s) from a DelegationHistory
func (k Keeper) AdjustDelegationTimestamps(delegationHistory types.DelegationHistory, difference sdk.Int) (types.DelegationHistory, error) {
	if !difference.IsNegative() {
		return delegationHistory, errors.New("difference has to be negative")
	}

	absoluteDifference := difference.Abs()

	for i := len(delegationHistory.GetHistory()) - 1; i >= 0; i-- {
		delegationTimestamp := delegationHistory.GetHistory()[i]

		if delegationTimestamp.GetBalance().Amount.GTE(absoluteDifference) {
			delegationHistory.History[i].Balance.Amount = delegationTimestamp.GetBalance().Amount.Sub(absoluteDifference)
			break
		} else {
			absoluteDifference = absoluteDifference.Sub(delegationTimestamp.GetBalance().Amount)
			delegationHistory.History[i].Balance.Amount = sdk.NewInt(0)
		}
	}

	return delegationHistory, nil
}

// PruneDelegationHistory prune the DelegationTimestamp history in a DelegationHistory
func (k Keeper) PruneDelegationHistory(delegationHistory types.DelegationHistory) types.DelegationHistory {
	delegationHistoryNew := types.DelegationHistory{}
	delegationHistoryNew.Address = delegationHistory.Address

	//remove any DelegationTimestamp that have a 0 amount
	for _, delegationTimestamp := range delegationHistory.GetHistory() {
		if !delegationTimestamp.GetBalance().Amount.Equal(sdk.NewInt(0)) {
			delegationHistoryNew.History = append(delegationHistoryNew.GetHistory(), delegationTimestamp)
		}
	}

	//TODO: add "staking length benefit period" to a param store
	//TODO: prune and compress if DelegationTimestamp is older than the "staking length benefit period"

	return delegationHistoryNew
}

// NewDelegationTimestamp creates a new DelegationTimestamp
func (k Keeper) NewDelegationTimestamp(ctx sdk.Context, amount sdk.Int) types.DelegationTimestamp {
	return types.DelegationTimestamp{
		Timestamp: time.Unix(ctx.BlockTime().Unix(), 0),
		Balance: sdk.NewCoin(
			k.stakingKeeper.BondDenom(ctx),
			amount,
		),
	}
}

// NewDelegationHistory creates a new DelegationHistory
func (k Keeper) NewDelegationHistory(ctx sdk.Context, delAddr sdk.AccAddress, delegatedAmount sdk.Int) types.DelegationHistory {
	delegationTimestamp := k.NewDelegationTimestamp(ctx, delegatedAmount)

	delegationHistory := types.DelegationHistory{
		Address: delAddr.String(),
		History: []*types.DelegationTimestamp{&delegationTimestamp},
	}

	return delegationHistory
}
