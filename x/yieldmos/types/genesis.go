package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DelegationHistoryList: []DelegationHistory{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in delegationHistory
	delegationHistoryIndexMap := make(map[string]struct{})

	for _, elem := range gs.DelegationHistoryList {
		index := string(DelegationHistoryKey(elem.Address))
		if _, ok := delegationHistoryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for delegationHistory")
		}
		delegationHistoryIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
