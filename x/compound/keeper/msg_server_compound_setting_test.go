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
	delegator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateCompoundSetting{Delegator: delegator,
			Index123: strconv.Itoa(i),
		}
		_, err := srv.CreateCompoundSetting(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetCompoundSetting(ctx,
			expected.Index123,
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
			desc: "Completed",
			request: &types.MsgUpdateCompoundSetting{Delegator: delegator,
				Index123: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateCompoundSetting{Delegator: "B",
				Index123: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateCompoundSetting{Delegator: delegator,
				Index123: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompoundKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateCompoundSetting{Delegator: delegator,
				Index123: strconv.Itoa(0),
			}
			_, err := srv.CreateCompoundSetting(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateCompoundSetting(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetCompoundSetting(ctx,
					expected.Index123,
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
			desc: "Completed",
			request: &types.MsgDeleteCompoundSetting{Delegator: delegator,
				Index123: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteCompoundSetting{Delegator: "B",
				Index123: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteCompoundSetting{Delegator: delegator,
				Index123: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CompoundKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateCompoundSetting(wctx, &types.MsgCreateCompoundSetting{Delegator: delegator,
				Index123: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteCompoundSetting(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetCompoundSetting(ctx,
					tc.request.Index123,
				)
				require.False(t, found)
			}
		})
	}
}
