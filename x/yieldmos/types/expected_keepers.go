package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingKeeper interface {
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (stakingTypes.Delegation, bool)
	BondDenom(ctx sdk.Context) string
	GetValidator(ctx sdk.Context, valAddr sdk.ValAddress) (stakingTypes.Validator, bool)
	GetAllDelegatorDelegations(ctx sdk.Context, delAddr sdk.AccAddress) []stakingTypes.Delegation
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}
