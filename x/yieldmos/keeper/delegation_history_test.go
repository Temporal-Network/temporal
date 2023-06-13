package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/testutil/nullify"
	"github.com/Temporal-Network/temporal/x/yieldmos/keeper"
	"github.com/Temporal-Network/temporal/x/yieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDelegationHistory(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DelegationHistory {
	items := make([]types.DelegationHistory, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetDelegationHistory(ctx, items[i])
	}
	return items
}

func TestDelegationHistoryGet(t *testing.T) {
	keeper, ctx := keepertest.YieldmosKeeper(t)
	items := createNDelegationHistory(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDelegationHistory(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDelegationHistoryRemove(t *testing.T) {
	keeper, ctx := keepertest.YieldmosKeeper(t)
	items := createNDelegationHistory(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDelegationHistory(ctx,
			item.Address,
		)
		_, found := keeper.GetDelegationHistory(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestDelegationHistoryGetAll(t *testing.T) {
	keeper, ctx := keepertest.YieldmosKeeper(t)
	items := createNDelegationHistory(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDelegationHistory(ctx)),
	)
}
