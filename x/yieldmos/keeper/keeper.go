package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/Temporal-Network/temporal/x/yieldmos/types"

	controllertypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper    types.AccountKeeper
		stakingKeeper    types.StakingKeeper
		msgServiceRouter *baseapp.MsgServiceRouter
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	stakingKeeper types.StakingKeeper,
	msgServiceRouter *baseapp.MsgServiceRouter,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		accountKeeper:    accountKeeper,
		stakingKeeper:    stakingKeeper,
		msgServiceRouter: msgServiceRouter,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) RegisterInterchainAccount(ctx sdk.Context) error {
	owner := k.accountKeeper.GetModuleAddress(types.ModuleName)
	msg := controllertypes.NewMsgRegisterInterchainAccount(
		"connection-0",
		owner.String(),
		"",
	)
	handler := k.msgServiceRouter.Handler(msg)
	res, err := handler(ctx, msg)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
