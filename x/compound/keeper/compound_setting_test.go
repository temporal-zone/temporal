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

func createNCompoundSetting(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CompoundSetting {
	items := make([]types.CompoundSetting, n)
	for i := range items {
		items[i].Delegator = strconv.Itoa(i)

		keeper.SetCompoundSetting(ctx, items[i])
	}
	return items
}

func TestCompoundSettingGet(t *testing.T) {
	keeper, ctx := keepertest.CompoundKeeper(t)
	items := createNCompoundSetting(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCompoundSetting(ctx,
			item.Delegator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCompoundSettingRemove(t *testing.T) {
	keeper, ctx := keepertest.CompoundKeeper(t)
	items := createNCompoundSetting(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCompoundSetting(ctx,
			item.Delegator,
		)
		_, found := keeper.GetCompoundSetting(ctx,
			item.Delegator,
		)
		require.False(t, found)
	}
}

func TestCompoundSettingGetAll(t *testing.T) {
	keeper, ctx := keepertest.CompoundKeeper(t)
	items := createNCompoundSetting(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCompoundSetting(ctx)),
	)
}
