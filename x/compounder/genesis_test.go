package compounder_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "temporal/testutil/keeper"
	"temporal/testutil/nullify"
	"temporal/x/compounder"
	"temporal/x/compounder/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CompoundSettingsList: []types.CompoundSettings{
			{
				Delegator: "0",
			},
			{
				Delegator: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CompounderKeeper(t)
	compounder.InitGenesis(ctx, *k, genesisState)
	got := compounder.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CompoundSettingsList, got.CompoundSettingsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
