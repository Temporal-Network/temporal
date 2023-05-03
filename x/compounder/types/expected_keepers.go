package types

import (
	"context"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	distrTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type DistrKeeper interface {
	DelegationTotalRewards(goCtx context.Context, req *distrTypes.QueryDelegationTotalRewardsRequest) (*distrTypes.QueryDelegationTotalRewardsResponse, error)
}

type StakingKeeper interface {
	Delegate(ctx sdk.Context, delAddr sdk.AccAddress, bondAmt math.Int, tokenSrc stakingTypes.BondStatus, validator stakingTypes.Validator, subtractAccount bool) (sdk.Dec, error)
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (stakingTypes.Validator, bool)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}
