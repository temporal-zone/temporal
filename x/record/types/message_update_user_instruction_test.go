package types

import (
	"testing"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/temporal-zone/temporal/testutil/sample"
)

func TestMsgUpdateUserInstruction_ValidateBasic(t *testing.T) {
	localAddress := sample.AccAddress()
	remoteAddress, err := DeriveRemoteAddress(localAddress, "juno")
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  MsgUpdateUserInstruction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateUserInstruction{
				LocalAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress",
			},
		}, {
			name: "invalid remote address",
			msg: MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   "remote",
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid freq",
			msg: MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       60,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "junoContractAddress",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid strat and contract",
			msg: MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       60,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      0,
				ContractAddress: "",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid strat and contract 1",
			msg: MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      1,
				ContractAddress: "",
			},
		}, {
			name: "valid strat and contract 2",
			msg: MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      0,
				ContractAddress: "junoContractAddress",
			},
		}, {
			name: "invalid instruction",
			msg: MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "juno-1",
				Frequency:       3600,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "",
				StrategyId:      0,
				ContractAddress: "junoContractAddress",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid chain id",
			msg: MsgUpdateUserInstruction{
				LocalAddress:    localAddress,
				RemoteAddress:   remoteAddress,
				ChainId:         "",
				Frequency:       3600,
				Expires:         time.Now().Add(time.Hour * 24),
				Instruction:     "{}",
				StrategyId:      0,
				ContractAddress: "junoContractAddress",
			},
			err: sdkerrors.ErrInvalidRequest,
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
