package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/record/keeper"
	"github.com/temporal-zone/temporal/x/record/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUserInstructions(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UserInstructions {
	items := make([]types.UserInstructions, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetUserInstructions(ctx, items[i])
	}
	return items
}

func TestUserInstructionsGet(t *testing.T) {
	k, ctx := keepertest.RecordKeeper(t)
	items := createNUserInstructions(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetUserInstructions(ctx, item.Address)
		require.True(t, found)
		require.Equal(t, nullify.Fill(&item), nullify.Fill(&rst))
	}
}
func TestUserInstructionsRemove(t *testing.T) {
	k, ctx := keepertest.RecordKeeper(t)
	items := createNUserInstructions(k, ctx, 10)
	for _, item := range items {
		k.RemoveUserInstructions(ctx, item.Address)
		_, found := k.GetUserInstructions(ctx, item.Address)
		require.False(t, found)
	}
}

func TestUserInstructionsGetAll(t *testing.T) {
	k, ctx := keepertest.RecordKeeper(t)
	items := createNUserInstructions(k, ctx, 10)
	require.ElementsMatch(t, nullify.Fill(items), nullify.Fill(k.GetAllUserInstructions(ctx)))
}
