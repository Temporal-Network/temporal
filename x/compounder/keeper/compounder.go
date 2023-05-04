package keeper

import (
	"cosmossdk.io/math"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	compTypes "temporal/x/compounder/types"
	"time"
)

// TODO: Change the amount of compoundings per block to be a module level param
const numberOfCompounds = 100

// TODO: RunCompounding needs test coverage
// RunCompounding gets all CompoundSettings and attempts to Compound them
func (k Keeper) RunCompounding(ctx sdk.Context) error {

	// TODO: Profile/Benchmark what 100, 1K, 10K, 100K of compounds in one block does to block time/other resource usage?
	numberOfCompoundsTemp := numberOfCompounds
	compSettings := k.GetAllCompoundSettings(ctx)

	for _, compSetting := range compSettings {
		shouldCompoundingHappen, err := k.ShouldCompoundingHappen(ctx, compSetting, ctx.BlockHeight(), ctx.BlockTime())
		if err != nil {
			return err
		}

		if !shouldCompoundingHappen {
			continue
		}

		err = k.Compound(ctx, compSetting)
		if err != nil {
			return err
		}

		numberOfCompoundsTemp--

		if numberOfCompoundsTemp <= 0 {
			return nil
		}
	}

	return nil
}

func (k Keeper) Compound(ctx sdk.Context, cs compTypes.CompoundSettings) error {
	address, err := sdk.AccAddressFromBech32(cs.Delegator)
	if err != nil {
		return err
	}

	// Get all active delegations
	req := distrTypes.QueryDelegationTotalRewardsRequest{DelegatorAddress: address.String()}
	delegations, err := k.distrKeeper.DelegationTotalRewards(ctx, &req)
	if err != nil {
		return err
	}

	// Calculate total rewards that can be claimed and delegated
	walletBalance := k.bankKeeper.GetBalance(ctx, address, sdk.DefaultBondDenom)
	amountToCompound := k.TotalCompoundAmount(delegations, walletBalance, cs)
	if amountToCompound.Amount.LT(sdk.NewInt(0)) {
		//TODO: Better error output that logs the delegations and compound settings
		return errors.New("amountToCompound is below 0 for: " + cs.Delegator)
	}

	if amountToCompound.Amount.Equal(sdk.NewInt(0)) {
		return nil
	}

	// Calculate each CompoundSettings validators compound amount
	totalCompoundPercent, compoundActions := k.BuildCompoundActions(cs, amountToCompound)
	if len(compoundActions) == 0 {
		return nil
	}

	if totalCompoundPercent.GT(sdk.NewInt(100)) {
		return errors.New("totalCompoundPercent can't be over 100")
	}

	// Handle any leftover amount if 100% of rewards are to be compounded by adding any leftover amount to their first validator
	compoundActions = k.HandleLeftOverAmount(compoundActions, totalCompoundPercent, amountToCompound)

	// Execute all CompoundActions
	for _, compoundAction := range compoundActions {
		err := Delegate(ctx, k, compoundAction, address)
		if err != nil {
			return err
		}
	}

	return nil
}

// ShouldCompoundingHappen compares the last time a compounding happened
func (k Keeper) ShouldCompoundingHappen(ctx sdk.Context, cs compTypes.CompoundSettings, blockHeight int64, blockTime time.Time) (bool, error) {
	if cs.Frequency == nil || cs.Frequency.IsNil() {
		return true, nil
	}

	previousCompounding, found := k.GetPreviousCompounding(ctx, cs.Delegator)
	if !found {
		return true, nil
	}

	if cs.Frequency.IsBlock() {
		currentBlockHeight := sdk.NewInt(blockHeight)
		nextCompoundHeight := previousCompounding.BlockHeight.Add(cs.Frequency.OnceEvery)

		return currentBlockHeight.GTE(nextCompoundHeight), nil
	}

	if cs.Frequency.IsSecond() {
		duration := time.Duration(cs.Frequency.OnceEvery.Int64()) * time.Second
		nextCompoundingTime := previousCompounding.Timestamp.Add(duration)

		return blockTime.After(nextCompoundingTime), nil
	}

	return true, nil
}

// Delegate is a helper method that delegates
func Delegate(ctx sdk.Context, k Keeper, compoundAction compTypes.StakingCompoundAction, address sdk.AccAddress) error {
	valAddr, err := sdk.ValAddressFromBech32(compoundAction.ValidatorAddress)
	if err != nil {
		return err
	}

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return errors.New("validator not found")
	}

	_, err = k.stakingKeeper.Delegate(ctx, address, compoundAction.Balance.Amount, stakingTypes.Unbonded, validator, true)
	if err != nil {
		return err
	}

	k.RecordCompounding(ctx, address.String(), ctx.BlockTime(), sdk.NewInt(ctx.BlockHeight()))

	return nil
}

// HandleLeftOverAmount calculates any leftover amount if totalCompoundPercent is 100%
func (k Keeper) HandleLeftOverAmount(compoundActions []compTypes.StakingCompoundAction, totalCompoundPercent math.Int, amountToCompound sdk.Coin) []compTypes.StakingCompoundAction {
	if totalCompoundPercent.Equal(sdk.NewInt(100)) {
		amountToBeCompounded := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(0))
		for _, compoundAmount := range compoundActions {
			amountToBeCompounded = amountToBeCompounded.Add(compoundAmount.Balance)
		}

		leftOver := amountToCompound.Amount.Sub(amountToBeCompounded.Amount)
		if leftOver.GT(sdk.NewInt(0)) {
			compoundActions[0].Balance.Amount = compoundActions[0].Balance.Amount.Add(leftOver)
		}
	}

	return compoundActions
}

// BuildCompoundActions creates the delegation actions that need to happen
func (k Keeper) BuildCompoundActions(cs compTypes.CompoundSettings, amountToCompound sdk.Coin) (math.Int, []compTypes.StakingCompoundAction) {
	totalCompoundPercent := sdk.NewInt(0)
	compoundActions := make([]compTypes.StakingCompoundAction, 0, 1)
	for _, valSetting := range cs.ValidatorSettings {
		validatorCompoundAmount := k.CalculateCompoundAmount(amountToCompound, valSetting.PercentToCompound)

		stakingCompoundAction := compTypes.StakingCompoundAction{
			ValidatorAddress: valSetting.ValidatorAddress,
			Balance:          sdk.NewCoin(amountToCompound.Denom, validatorCompoundAmount),
		}

		compoundActions = append(compoundActions, stakingCompoundAction)

		totalCompoundPercent = totalCompoundPercent.Add(valSetting.PercentToCompound)
	}

	return totalCompoundPercent, compoundActions
}

// TotalCompoundAmount sums all delegations and extra balance amount
func (k Keeper) TotalCompoundAmount(delegations *distrTypes.QueryDelegationTotalRewardsResponse, walletBalance sdk.Coin, cs compTypes.CompoundSettings) sdk.Coin {
	// Sum the total staking claims
	outstandingRewards := k.StakingCompoundAmount(delegations, walletBalance)

	// Extra balance above CompoundSettings.AmountToRemain
	extraCompoundAmount := k.ExtraCompoundAmount(cs, walletBalance)

	return outstandingRewards.Add(extraCompoundAmount)
}

func (k Keeper) StakingCompoundAmount(delegations *distrTypes.QueryDelegationTotalRewardsResponse, walletBalance sdk.Coin) sdk.Coin {
	outstandingRewards := sdk.Coin{Denom: walletBalance.Denom, Amount: sdk.NewInt(0)}
	for _, delegation := range delegations.GetRewards() {
		for _, reward := range delegation.Reward {
			outstandingRewards = outstandingRewards.AddAmount(reward.Amount.TruncateInt())
		}
	}

	return outstandingRewards
}

// ExtraCompoundAmount calcs the diff between CompoundSettings.AmountToRemain and the wallet balance
func (k Keeper) ExtraCompoundAmount(cs compTypes.CompoundSettings, walletBalance sdk.Coin) sdk.Coin {
	extraCompoundAmount := sdk.Coin{Denom: walletBalance.Denom, Amount: sdk.NewInt(0)}

	if cs.AmountToRemain.IsNil() {
		return extraCompoundAmount
	}

	if walletBalance.Denom == cs.AmountToRemain.Denom && walletBalance.Amount.GT(cs.AmountToRemain.Amount) {
		extraCompoundAmount = walletBalance.Sub(cs.AmountToRemain)
	}

	return extraCompoundAmount
}

// RecordCompounding records the compounding timestamp
func (k Keeper) RecordCompounding(ctx sdk.Context, address string, blockTime time.Time, blockHeight math.Int) {
	value, _ := k.GetPreviousCompounding(ctx, address)

	value.Delegator = address
	value.Timestamp = blockTime
	value.BlockHeight = blockHeight

	k.SetPreviousCompounding(ctx, value)
}

// CalculateCompoundAmount calcs the compounding amount
func (k Keeper) CalculateCompoundAmount(rewardAmount sdk.Coin, percentToCompound math.Int) math.Int {
	amountToCompound := rewardAmount.Amount.Mul(percentToCompound).Quo(sdk.NewInt(100))

	return amountToCompound
}
