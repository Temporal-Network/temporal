package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Hooks wrapper struct for stakestore keeper
type Hooks struct {
	k Keeper
}

var _ stakingtypes.StakingHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) AfterValidatorCreated(sdk.Context, sdk.ValAddress) error                  { return nil }
func (h Hooks) BeforeValidatorModified(sdk.Context, sdk.ValAddress) error                { return nil }
func (h Hooks) AfterValidatorRemoved(sdk.Context, sdk.ConsAddress, sdk.ValAddress) error { return nil }
func (h Hooks) AfterValidatorBonded(sdk.Context, sdk.ConsAddress, sdk.ValAddress) error  { return nil }
func (h Hooks) AfterValidatorBeginUnbonding(sdk.Context, sdk.ConsAddress, sdk.ValAddress) error {
	return nil
}
func (h Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}
func (h Hooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}
func (h Hooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	err := h.k.CheckDelegationHistoryRecords(ctx, delAddr)
	if err != nil {
		return err
	}

	return nil
}

func (h Hooks) BeforeValidatorSlashed(sdk.Context, sdk.ValAddress, sdk.Dec) error { return nil }
