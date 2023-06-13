package compounder_test

import (
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/testutil/nullify"
	"github.com/Temporal-Network/temporal/x/compounder"
	"github.com/Temporal-Network/temporal/x/compounder/types"
	"github.com/stretchr/testify/require"
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
		PreviousCompoundingList: []types.PreviousCompounding{
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
	require.ElementsMatch(t, genesisState.PreviousCompoundingList, got.PreviousCompoundingList)
	// this line is used by starport scaffolding # genesis/test/assert
}
