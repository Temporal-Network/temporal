package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated contractRemoteZone",
			genState: &types.GenesisState{
				ContractRemoteZoneList: []types.ContractRemoteZone{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid contractRemoteZone count",
			genState: &types.GenesisState{
				ContractRemoteZoneList: []types.ContractRemoteZone{
					{
						Id: 1,
					},
				},
				ContractRemoteZoneCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated remoteContractCompoundSettings",
			genState: &types.GenesisState{
				RemoteContractCompoundSettingsList: []types.RemoteContractCompoundSettings{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid remoteContractCompoundSettings count",
			genState: &types.GenesisState{
				RemoteContractCompoundSettingsList: []types.RemoteContractCompoundSettings{
					{
						Id: 1,
					},
				},
				RemoteContractCompoundSettingsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated previousRemoteCompounding",
			genState: &types.GenesisState{
				PreviousRemoteCompoundingList: []types.PreviousRemoteCompounding{
					{
						RemoteContractCompoundSetting: 0,
					},
					{
						RemoteContractCompoundSetting: 0,
					},
				},
			},
			valid: false,
		},
		{
	desc:     "duplicated iCARemoteZone",
	genState: &types.GenesisState{
		ICARemoteZoneList: []types.ICARemoteZone{
			{
				Id: 0,
			},
			{
				Id: 0,
			},
		},
	},
	valid:    false,
},
{
	desc:     "invalid iCARemoteZone count",
	genState: &types.GenesisState{
		ICARemoteZoneList: []types.ICARemoteZone{
			{
				Id: 1,
			},
		},
		ICARemoteZoneCount: 0,
	},
	valid:    false,
},
// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
