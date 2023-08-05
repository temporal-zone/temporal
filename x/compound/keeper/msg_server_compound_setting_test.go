package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/x/compound/keeper"
	"github.com/temporal-zone/temporal/x/compound/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCompoundSettingMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CompoundKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	for i := 0; i < 5; i++ {
		delegator := strconv.Itoa(i)
		expected := &types.MsgCreateCompoundSetting{Delegator: delegator}
		_, err := srv.CreateCompoundSetting(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetCompoundSetting(ctx,
			expected.Delegator,
		)
		require.True(t, found)
		require.Equal(t, expected.Delegator, rst.Delegator)
	}
}

func TestCompoundSettingMsgServerUpdate(t *testing.T) {
	delegator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateCompoundSetting
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateCompoundSetting{Delegator: delegator},
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgUpdateCompoundSetting{Delegator: delegator},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompoundKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateCompoundSetting{Delegator: delegator}
			_, err := srv.CreateCompoundSetting(wctx, expected)
			require.NoError(t, err)

			if tc.desc == "KeyNotFound" {
				delete := &types.MsgDeleteCompoundSetting{Delegator: delegator}
				_, err = srv.DeleteCompoundSetting(wctx, delete)
				require.NoError(t, err)
			}

			_, err = srv.UpdateCompoundSetting(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetCompoundSetting(ctx,
					expected.Delegator,
				)
				require.True(t, found)
				require.Equal(t, expected.Delegator, rst.Delegator)
			}
		})
	}
}

func TestCompoundSettingMsgServerDelete(t *testing.T) {
	delegator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteCompoundSetting
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteCompoundSetting{Delegator: delegator},
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteCompoundSetting{Delegator: delegator},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompoundKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateCompoundSetting(wctx, &types.MsgCreateCompoundSetting{Delegator: delegator})
			require.NoError(t, err)

			if tc.desc == "KeyNotFound" {
				delete := &types.MsgDeleteCompoundSetting{Delegator: delegator}
				_, err := srv.DeleteCompoundSetting(wctx, delete)
				require.NoError(t, err)
			}

			_, err = srv.DeleteCompoundSetting(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetCompoundSetting(ctx,
					tc.request.Delegator,
				)
				require.False(t, found)
			}
		})
	}
}

// TODO when minimumCompoundFrequency becomes a module level param, change this tests uint64(600) to test against the param.
func TestCheckFrequency(t *testing.T) {
	tests := []struct {
		desc    string
		request uint64
		err     error
	}{
		{
			desc:    "Under",
			request: 10,
		},
		{
			desc:    "Over",
			request: 700,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, _ := keepertest.CompoundKeeper(t)

			onceEvery := k.CheckFrequency(tc.request)

			if tc.desc == "Under" {
				require.GreaterOrEqual(t, onceEvery, uint64(600))
			}

			if tc.desc == "Over" {
				require.Equal(t, onceEvery, tc.request)
			}
		})
	}
}
