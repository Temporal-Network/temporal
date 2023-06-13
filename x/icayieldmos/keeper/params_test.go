package keeper_test

import (
	"testing"

	testkeeper "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IcayieldmosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
