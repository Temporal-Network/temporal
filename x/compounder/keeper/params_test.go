package keeper_test

import (
	"testing"

	testkeeper "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/x/compounder/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CompounderKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
