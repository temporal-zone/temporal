package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgCreateUserInstruction = "create_user_instruction"

var _ sdk.Msg = &MsgCreateUserInstruction{}

func NewMsgCreateUserInstruction(
	localAddress string,
	remoteAddress string,
	chainId string,
	frequency int64,
	expires time.Time,
	instruction string,
	strategyId int64,
	contractAddress string) *MsgCreateUserInstruction {
	return &MsgCreateUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   remoteAddress,
		ChainId:         chainId,
		Frequency:       frequency,
		Expires:         expires,
		Instruction:     instruction,
		StrategyId:      strategyId,
		ContractAddress: contractAddress,
	}
}

func (msg *MsgCreateUserInstruction) Route() string {
	return RouterKey
}

func (msg *MsgCreateUserInstruction) Type() string {
	return TypeMsgCreateUserInstruction
}

func (msg *MsgCreateUserInstruction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.LocalAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateUserInstruction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUserInstruction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.LocalAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
