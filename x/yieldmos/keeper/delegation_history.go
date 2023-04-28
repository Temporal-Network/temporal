package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"temporal/x/yieldmos/types"
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
