package types_test

import (
	"testing"

	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"github.com/stretchr/testify/require"
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
