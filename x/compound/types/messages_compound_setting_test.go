package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"temporal/testutil/sample"
)

func TestMsgCreateCompoundSetting_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCompoundSetting
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCompoundSetting{
				Delegator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCompoundSetting{
				Delegator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateCompoundSetting_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCompoundSetting
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCompoundSetting{
				Delegator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCompoundSetting{
				Delegator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteCompoundSetting_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteCompoundSetting
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteCompoundSetting{
				Delegator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteCompoundSetting{
				Delegator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
