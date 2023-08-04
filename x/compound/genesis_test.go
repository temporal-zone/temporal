package compound_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "temporal/testutil/keeper"
	"temporal/testutil/nullify"
	"temporal/x/compound"
	"temporal/x/compound/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CompoundKeeper(t)
	compound.InitGenesis(ctx, *k, genesisState)
	got := compound.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CompoundSettingList, got.CompoundSettingList)
	require.ElementsMatch(t, genesisState.PreviousCompoundingList, got.PreviousCompoundingList)
	// this line is used by starport scaffolding # genesis/test/assert
}
