package keeper_test

import (
	"testing"

    "github.com/temporal-zone/temporal/x/icayieldmos/keeper"
    "github.com/temporal-zone/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/stretchr/testify/require"
)

func createNICARemoteZone(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ICARemoteZone {
	items := make([]types.ICARemoteZone, n)
	for i := range items {
		items[i].Id = keeper.AppendICARemoteZone(ctx, items[i])
	}
	return items
}

func TestICARemoteZoneGet(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNICARemoteZone(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetICARemoteZone(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestICARemoteZoneRemove(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNICARemoteZone(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveICARemoteZone(ctx, item.Id)
		_, found := keeper.GetICARemoteZone(ctx, item.Id)
		require.False(t, found)
	}
}

func TestICARemoteZoneGetAll(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNICARemoteZone(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllICARemoteZone(ctx)),
	)
}

func TestICARemoteZoneCount(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNICARemoteZone(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetICARemoteZoneCount(ctx))
}
