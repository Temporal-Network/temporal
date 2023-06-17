package keeper_test

import (
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/testutil/nullify"
	"github.com/Temporal-Network/temporal/x/icayieldmos/keeper"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNContractRemoteZone(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ContractRemoteZone {
	items := make([]types.ContractRemoteZone, n)
	for i := range items {
		items[i].Id = keeper.AppendContractRemoteZone(ctx, items[i])
	}
	return items
}

func TestContractRemoteZoneGet(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNContractRemoteZone(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetContractRemoteZone(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestContractRemoteZoneRemove(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNContractRemoteZone(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveContractRemoteZone(ctx, item.Id)
		_, found := keeper.GetContractRemoteZone(ctx, item.Id)
		require.False(t, found)
	}
}

func TestContractRemoteZoneGetAll(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNContractRemoteZone(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllContractRemoteZone(ctx)),
	)
}

func TestContractRemoteZoneCount(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNContractRemoteZone(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetContractRemoteZoneCount(ctx))
}
