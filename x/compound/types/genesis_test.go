package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"temporal/x/compound/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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

				CompoundSettingList: []types.CompoundSetting{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated compoundSetting",
			genState: &types.GenesisState{
				CompoundSettingList: []types.CompoundSetting{
					{
						Delegator: "0",
					},
					{
						Delegator: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated previousCompounding",
			genState: &types.GenesisState{
				PreviousCompoundingList: []types.PreviousCompounding{
					{
						Delegator: "0",
					},
					{
						Delegator: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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
