package yieldmos_test

import (
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/testutil/nullify"
	"github.com/Temporal-Network/temporal/x/yieldmos"
	"github.com/Temporal-Network/temporal/x/yieldmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DelegationHistoryList: []types.DelegationHistory{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.YieldmosKeeper(t)
	yieldmos.InitGenesis(ctx, *k, genesisState)
	got := yieldmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DelegationHistoryList, got.DelegationHistoryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
