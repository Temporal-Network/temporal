package icayieldmos_test

import (
	"testing"

	keepertest "github.com/Temporal-Network/temporal/testutil/keeper"
	"github.com/Temporal-Network/temporal/testutil/nullify"
	"github.com/Temporal-Network/temporal/x/icayieldmos"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		ContractRemoteZoneList: []types.ContractRemoteZone{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ContractRemoteZoneCount: 2,
		RemoteContractCompoundSettingsList: []types.RemoteContractCompoundSettings{
		{
			Id: 0,
		},
		{
			Id: 1,
		},
	},
	RemoteContractCompoundSettingsCount: 2,
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IcayieldmosKeeper(t)
	icayieldmos.InitGenesis(ctx, *k, genesisState)
	got := icayieldmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.ContractRemoteZoneList, got.ContractRemoteZoneList)
	require.Equal(t, genesisState.ContractRemoteZoneCount, got.ContractRemoteZoneCount)
	require.ElementsMatch(t, genesisState.RemoteContractCompoundSettingsList, got.RemoteContractCompoundSettingsList)
require.Equal(t, genesisState.RemoteContractCompoundSettingsCount, got.RemoteContractCompoundSettingsCount)
// this line is used by starport scaffolding # genesis/test/assert
}
