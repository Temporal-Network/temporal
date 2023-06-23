package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/x/compounder/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CompounderKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
