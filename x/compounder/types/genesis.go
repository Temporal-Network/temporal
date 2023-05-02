package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CompoundSettingsList: []CompoundSettings{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in compoundSettings
	compoundSettingsIndexMap := make(map[string]struct{})

	for _, elem := range gs.CompoundSettingsList {
		index := string(CompoundSettingsKey(elem.Delegator))
		if _, ok := compoundSettingsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for compoundSettings")
		}
		compoundSettingsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
