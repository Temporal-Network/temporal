package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/testutil/nullify"
	"github.com/Temporal-Network/temporal/x/compounder/keeper"
	"github.com/Temporal-Network/temporal/x/compounder/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCompoundSettings(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CompoundSettings {
	items := make([]types.CompoundSettings, n)
	for i := range items {
		items[i].Delegator = strconv.Itoa(i)

		keeper.SetCompoundSettings(ctx, items[i])
	}
	return items
}

func TestCompoundSettingsGet(t *testing.T) {
	keeper, ctx := keepertest.CompounderKeeper(t)
	items := createNCompoundSettings(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCompoundSettings(ctx,
			item.Delegator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCompoundSettingsRemove(t *testing.T) {
	keeper, ctx := keepertest.CompounderKeeper(t)
	items := createNCompoundSettings(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCompoundSettings(ctx,
			item.Delegator,
		)
		_, found := keeper.GetCompoundSettings(ctx,
			item.Delegator,
		)
		require.False(t, found)
	}
}

func TestCompoundSettingsGetAll(t *testing.T) {
	keeper, ctx := keepertest.CompounderKeeper(t)
	items := createNCompoundSettings(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCompoundSettings(ctx)),
	)
}
