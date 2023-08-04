package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "temporal/testutil/keeper"
	"temporal/x/compound/keeper"
	"temporal/x/compound/types"
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
