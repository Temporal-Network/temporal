package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/x/compounder/keeper"
	"github.com/Temporal-Network/temporal/x/compounder/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CompounderKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
