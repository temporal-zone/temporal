package keeper_test

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/temporal-zone/temporal/app/apptesting"
	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/record/keeper"
	"github.com/temporal-zone/temporal/x/record/types"
	"strconv"
	"testing"
	"time"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDelegationHistory(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DelegationHistory {
	items := make([]types.DelegationHistory, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetDelegationHistory(ctx, items[i])
	}
	return items
}

func TestDelegationHistoryGet(t *testing.T) {
	k, ctx := keepertest.RecordKeeper(t)
	items := createNDelegationHistory(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetDelegationHistory(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDelegationHistoryRemove(t *testing.T) {
	k, ctx := keepertest.RecordKeeper(t)
	items := createNDelegationHistory(k, ctx, 10)
	for _, item := range items {
		k.RemoveDelegationHistory(ctx,
			item.Address,
		)
		_, found := k.GetDelegationHistory(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestDelegationHistoryGetAll(t *testing.T) {
	k, ctx := keepertest.RecordKeeper(t)
	items := createNDelegationHistory(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllDelegationHistory(ctx)),
	)
}

func TestDelegationTimestamps(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	var tests = []struct {
		desc            string
		startingAmount  math.Int
		differences     []math.Int
		expectedAmounts []math.Int
		expectedHistory int
	}{
		{
			desc:            "Single Positive",
			startingAmount:  math.NewInt(0),
			differences:     []math.Int{math.NewInt(500)},
			expectedAmounts: []math.Int{math.NewInt(500)},
			expectedHistory: 1,
		},
		{
			desc:            "Single Negative",
			startingAmount:  math.NewInt(1000),
			differences:     []math.Int{math.NewInt(-500)},
			expectedAmounts: []math.Int{math.NewInt(500)},
			expectedHistory: 1,
		},
		{
			desc:            "Double Positive",
			startingAmount:  math.NewInt(1000),
			differences:     []math.Int{math.NewInt(500), math.NewInt(500)},
			expectedAmounts: []math.Int{math.NewInt(1500), math.NewInt(2000)},
			expectedHistory: 1,
		},
		{
			desc:            "Double Negative",
			startingAmount:  math.NewInt(1000),
			differences:     []math.Int{math.NewInt(-500), math.NewInt(-500)},
			expectedAmounts: []math.Int{math.NewInt(500), math.NewInt(0)},
			expectedHistory: 0,
		},
		{
			desc:            "Double Mixed 1",
			startingAmount:  math.NewInt(1000),
			differences:     []math.Int{math.NewInt(500), math.NewInt(-500)},
			expectedAmounts: []math.Int{math.NewInt(1500), math.NewInt(1000)},
			expectedHistory: 1,
		},
		{
			desc:            "Double Mixed 2",
			startingAmount:  math.NewInt(1000),
			differences:     []math.Int{math.NewInt(-500), math.NewInt(500)},
			expectedAmounts: []math.Int{math.NewInt(500), math.NewInt(1000)},
			expectedHistory: 1,
		},
		{
			desc:            "Double Long Time 1",
			startingAmount:  math.NewInt(1000),
			differences:     []math.Int{math.NewInt(500), math.NewInt(500)},
			expectedAmounts: []math.Int{math.NewInt(1500), math.NewInt(2000)},
			expectedHistory: 1,
		},
	}

	for _, tc := range tests {
		require.True(t, len(tc.differences) == len(tc.expectedAmounts))

		delegationTimestamp := s.App.RecordKeeper.NewDelegationTimestamp(s.Ctx, tc.startingAmount)
		delegationHistory := types.DelegationHistory{
			Address: "address1",
			History: []*types.DelegationTimestamp{&delegationTimestamp},
		}

		for i := 0; i < len(tc.differences); i++ {
			var adjustedDelHistory types.DelegationHistory
			if tc.differences[i].IsPositive() {
				adjustedDelHistory = s.App.RecordKeeper.AddDelegationTimestamp(s.Ctx, delegationHistory, tc.differences[i])
			}

			if tc.differences[i].IsNegative() {
				adjustedDelHistory = s.App.RecordKeeper.RemoveDelegationTimestamps(delegationHistory, tc.differences[i])
			}

			if tc.expectedHistory > 0 {
				require.Equal(t, len(adjustedDelHistory.GetHistory()), tc.expectedHistory)
				require.Equal(t, adjustedDelHistory.GetHistory()[0].Balance.Amount, tc.expectedAmounts[i])
			}
		}
	}
}

func TestDelegationTimestampsMultiDay(t *testing.T) {
	s := apptesting.SetupSuitelessTestHelper()

	var tests = []struct {
		desc            string
		startingAmount  math.Int
		differences     []math.Int
		expectedAmounts []math.Int
		numExpDaily     []int
		numDays         int
	}{
		{
			desc:            "Two Days Additive",
			startingAmount:  math.NewInt(1000),
			differences:     []math.Int{math.NewInt(500), math.NewInt(500)},
			expectedAmounts: []math.Int{math.NewInt(1500), math.NewInt(2000)},
			numExpDaily:     []int{2, 3},
			numDays:         2,
		},
		{
			desc:            "Two Days Reductive",
			startingAmount:  math.NewInt(1000),
			differences:     []math.Int{math.NewInt(-500), math.NewInt(-500)},
			expectedAmounts: []math.Int{math.NewInt(500), math.NewInt(0)},
			numExpDaily:     []int{1, 0},
			numDays:         2,
		},
	}

	for _, tc := range tests {
		require.True(t, len(tc.differences) == len(tc.expectedAmounts))

		delegationTimestamp := s.App.RecordKeeper.NewDelegationTimestamp(s.Ctx, tc.startingAmount)
		delegationHistory := types.DelegationHistory{
			Address: "address1",
			History: []*types.DelegationTimestamp{&delegationTimestamp},
		}

		ctx := s.Ctx.WithBlockTime(s.Ctx.BlockTime().Add(time.Hour * 24))

		for i := 0; i < tc.numDays; i++ {
			fmt.Println(ctx.BlockTime())

			var adjustedDelHistory types.DelegationHistory

			if tc.differences[i].IsPositive() {
				adjustedDelHistory = s.App.RecordKeeper.AddDelegationTimestamp(ctx, delegationHistory, tc.differences[i])
			}

			if tc.differences[i].IsNegative() {
				adjustedDelHistory = s.App.RecordKeeper.RemoveDelegationTimestamps(delegationHistory, tc.differences[i])
			}

			require.Equal(t, len(adjustedDelHistory.GetHistory()), tc.numExpDaily[i])

			totalAmount := math.NewInt(0)
			for _, delTs := range adjustedDelHistory.GetHistory() {
				totalAmount = totalAmount.Add(delTs.Balance.Amount)
			}

			require.True(t, totalAmount.Equal(tc.expectedAmounts[i]))

			ctx = s.Ctx.WithBlockTime(ctx.BlockTime().Add(time.Hour * 24))
			delegationHistory = adjustedDelHistory
		}
	}
}
