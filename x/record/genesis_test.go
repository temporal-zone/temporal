package record_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/record"
	"github.com/temporal-zone/temporal/x/record/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DelegationHistoryList: []types.DelegationHistory{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RecordKeeper(t)
	record.InitGenesis(ctx, *k, genesisState)
	got := record.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DelegationHistoryList, got.DelegationHistoryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
