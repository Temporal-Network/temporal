package types

import (
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		ContractRemoteZoneList: []ContractRemoteZone{},
// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in contractRemoteZone
contractRemoteZoneIdMap := make(map[uint64]bool)
contractRemoteZoneCount := gs.GetContractRemoteZoneCount()
for _, elem := range gs.ContractRemoteZoneList {
	if _, ok := contractRemoteZoneIdMap[elem.Id]; ok {
		return fmt.Errorf("duplicated id for contractRemoteZone")
	}
	if elem.Id >= contractRemoteZoneCount {
		return fmt.Errorf("contractRemoteZone id should be lower or equal than the last id")
	}
	contractRemoteZoneIdMap[elem.Id] = true
}
// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
