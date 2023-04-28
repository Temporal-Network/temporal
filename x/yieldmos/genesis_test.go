package yieldmos_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "temporal/testutil/keeper"
	"temporal/testutil/nullify"
	"temporal/x/yieldmos"
	"temporal/x/yieldmos/types"
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
