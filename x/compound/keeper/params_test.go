package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/x/compound/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CompoundKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.NumberOfCompoundsPerBlock, k.NumberOfCompoundsPerBlock(ctx))
	require.EqualValues(t, params.MinimumCompoundFrequency, k.MinimumCompoundFrequency(ctx))
	require.EqualValues(t, params.CompoundModuleEnabled, k.CompoundModuleEnabled(ctx))
}
