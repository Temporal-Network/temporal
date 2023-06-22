package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:                 PortID,
		ContractRemoteZoneList: []ContractRemoteZone{},
		RemoteContractCompoundSettingsList: []RemoteContractCompoundSettings{},
PreviousRemoteCompoundingList: []PreviousRemoteCompounding{},
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
	// Check for duplicated ID in remoteContractCompoundSettings
remoteContractCompoundSettingsIdMap := make(map[uint64]bool)
remoteContractCompoundSettingsCount := gs.GetRemoteContractCompoundSettingsCount()
for _, elem := range gs.RemoteContractCompoundSettingsList {
	if _, ok := remoteContractCompoundSettingsIdMap[elem.Id]; ok {
		return fmt.Errorf("duplicated id for remoteContractCompoundSettings")
	}
	if elem.Id >= remoteContractCompoundSettingsCount {
		return fmt.Errorf("remoteContractCompoundSettings id should be lower or equal than the last id")
	}
	remoteContractCompoundSettingsIdMap[elem.Id] = true
}
// Check for duplicated index in previousRemoteCompounding
previousRemoteCompoundingIndexMap := make(map[string]struct{})

for _, elem := range gs.PreviousRemoteCompoundingList {
	index := string(PreviousRemoteCompoundingKey(elem.RemoteContractCompoundSetting))
	if _, ok := previousRemoteCompoundingIndexMap[index]; ok {
		return fmt.Errorf("duplicated index for previousRemoteCompounding")
	}
	previousRemoteCompoundingIndexMap[index] = struct{}{}
}
// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
