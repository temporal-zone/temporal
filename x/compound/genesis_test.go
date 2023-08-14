package compound_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/compound"
	"github.com/temporal-zone/temporal/x/compound/types"
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
		PreviousCompoundList: []types.PreviousCompound{
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
	require.ElementsMatch(t, genesisState.PreviousCompoundList, got.PreviousCompoundList)
	// this line is used by starport scaffolding # genesis/test/assert
}
