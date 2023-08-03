package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "temporal/testutil/keeper"
	"temporal/testutil/nullify"
	"temporal/x/compound/keeper"
	"temporal/x/compound/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCompoundSetting(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CompoundSetting {
	items := make([]types.CompoundSetting, n)
	for i := range items {
		items[i].Index123 = strconv.Itoa(i)

		keeper.SetCompoundSetting(ctx, items[i])
	}
	return items
}

func TestCompoundSettingGet(t *testing.T) {
	keeper, ctx := keepertest.CompoundKeeper(t)
	items := createNCompoundSetting(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCompoundSetting(ctx,
			item.Index123,
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
			item.Index123,
		)
		_, found := keeper.GetCompoundSetting(ctx,
			item.Index123,
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
