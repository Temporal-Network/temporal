package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/testutil/nullify"
	"github.com/Temporal-Network/temporal/x/icayieldmos/keeper"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPreviousRemoteCompounding(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PreviousRemoteCompounding {
	items := make([]types.PreviousRemoteCompounding, n)
	for i := range items {
		items[i].RemoteContractCompoundSetting = uint64(i)

		keeper.SetPreviousRemoteCompounding(ctx, items[i])
	}
	return items
}

func TestPreviousRemoteCompoundingGet(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNPreviousRemoteCompounding(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPreviousRemoteCompounding(ctx,
			item.RemoteContractCompoundSetting,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPreviousRemoteCompoundingRemove(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNPreviousRemoteCompounding(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePreviousRemoteCompounding(ctx,
			item.RemoteContractCompoundSetting,
		)
		_, found := keeper.GetPreviousRemoteCompounding(ctx,
			item.RemoteContractCompoundSetting,
		)
		require.False(t, found)
	}
}

func TestPreviousRemoteCompoundingGetAll(t *testing.T) {
	keeper, ctx := keepertest.IcayieldmosKeeper(t)
	items := createNPreviousRemoteCompounding(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPreviousRemoteCompounding(ctx)),
	)
}
