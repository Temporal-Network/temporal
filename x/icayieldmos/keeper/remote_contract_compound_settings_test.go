package keeper_test

import (
	"testing"

    "github.com/Temporal-Network/temporal/x/icayieldmos/keeper"
    "github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/testutil/nullify"
	"github.com/stretchr/testify/require"
)

func createNRemoteContractCompoundSettings(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RemoteContractCompoundSettings {
	items := make([]types.RemoteContractCompoundSettings, n)
	for i := range items {
		items[i].Id = keeper.AppendRemoteContractCompoundSettings(ctx, items[i])
	}
	return items
}

func TestRemoteContractCompoundSettingsGet(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNRemoteContractCompoundSettings(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRemoteContractCompoundSettings(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRemoteContractCompoundSettingsRemove(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNRemoteContractCompoundSettings(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRemoteContractCompoundSettings(ctx, item.Id)
		_, found := keeper.GetRemoteContractCompoundSettings(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRemoteContractCompoundSettingsGetAll(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNRemoteContractCompoundSettings(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRemoteContractCompoundSettings(ctx)),
	)
}

func TestRemoteContractCompoundSettingsCount(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNRemoteContractCompoundSettings(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRemoteContractCompoundSettingsCount(ctx))
}
