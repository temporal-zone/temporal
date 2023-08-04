package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CompoundSettingList: []CompoundSetting{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in compoundSetting
	compoundSettingIndexMap := make(map[string]struct{})

	for _, elem := range gs.CompoundSettingList {
		index := string(CompoundSettingKey(elem.Delegator))
		if _, ok := compoundSettingIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for compoundSetting")
		}
		compoundSettingIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
