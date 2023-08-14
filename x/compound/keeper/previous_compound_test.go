package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/compound/keeper"
	"github.com/temporal-zone/temporal/x/compound/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPreviousCompound(compKeeper *keeper.Keeper, ctx sdk.Context, n int) []types.PreviousCompound {
	items := make([]types.PreviousCompound, n)
	for i := range items {
		items[i].Delegator = strconv.Itoa(i)

		compKeeper.SetPreviousCompound(ctx, items[i])
	}
	return items
}

func TestPreviousCompoundGet(t *testing.T) {
	compKeeper, ctx := keepertest.CompoundKeeper(t)
	items := createNPreviousCompound(compKeeper, ctx, 10)
	for _, item := range items {
		rst, found := compKeeper.GetPreviousCompound(ctx,
			item.Delegator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPreviousCompoundRemove(t *testing.T) {
	compKeeper, ctx := keepertest.CompoundKeeper(t)
	items := createNPreviousCompound(compKeeper, ctx, 10)
	for _, item := range items {
		compKeeper.RemovePreviousCompound(ctx,
			item.Delegator,
		)
		_, found := compKeeper.GetPreviousCompound(ctx,
			item.Delegator,
		)
		require.False(t, found)
	}
}

func TestPreviousCompoundGetAll(t *testing.T) {
	compKeeper, ctx := keepertest.CompoundKeeper(t)
	items := createNPreviousCompound(compKeeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(compKeeper.GetAllPreviousCompound(ctx)),
	)
}
