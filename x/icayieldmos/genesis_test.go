package icayieldmos_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/icayieldmos"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
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
		PreviousRemoteCompoundingList: []types.PreviousRemoteCompounding{
			{
				RemoteContractCompoundSetting: 0,
			},
			{
				RemoteContractCompoundSetting: 1,
			},
		},
		ICARemoteZoneList: []types.ICARemoteZone{
		{
			Id: 0,
		},
		{
			Id: 1,
		},
	},
	ICARemoteZoneCount: 2,
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
	require.ElementsMatch(t, genesisState.PreviousRemoteCompoundingList, got.PreviousRemoteCompoundingList)
	require.ElementsMatch(t, genesisState.ICARemoteZoneList, got.ICARemoteZoneList)
require.Equal(t, genesisState.ICARemoteZoneCount, got.ICARemoteZoneCount)
// this line is used by starport scaffolding # genesis/test/assert
}
