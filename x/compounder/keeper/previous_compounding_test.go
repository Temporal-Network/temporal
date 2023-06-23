package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/compounder/keeper"
	"github.com/temporal-zone/temporal/x/compounder/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPreviousCompounding(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PreviousCompounding {
	items := make([]types.PreviousCompounding, n)
	for i := range items {
		items[i].Delegator = strconv.Itoa(i)
		items[i].BlockHeight = sdk.NewInt(int64(i))

		keeper.SetPreviousCompounding(ctx, items[i])
	}
	return items
}

func TestPreviousCompoundingGet(t *testing.T) {
	keeper, ctx := keepertest.CompounderKeeper(t)
	items := createNPreviousCompounding(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPreviousCompounding(ctx,
			item.Delegator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPreviousCompoundingRemove(t *testing.T) {
	keeper, ctx := keepertest.CompounderKeeper(t)
	items := createNPreviousCompounding(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePreviousCompounding(ctx,
			item.Delegator,
		)
		_, found := keeper.GetPreviousCompounding(ctx,
			item.Delegator,
		)
		require.False(t, found)
	}
}

func TestPreviousCompoundingGetAll(t *testing.T) {
	keeper, ctx := keepertest.CompounderKeeper(t)
	items := createNPreviousCompounding(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPreviousCompounding(ctx)),
	)
}
