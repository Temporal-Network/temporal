package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/x/icayieldmos/keeper"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IcayieldmosKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
