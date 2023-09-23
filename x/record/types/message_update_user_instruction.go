package types

import (
	sdkerr "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgUpdateUserInstruction = "update_user_instruction"

var _ sdk.Msg = &MsgUpdateUserInstruction{}

func NewMsgUpdateUserInstruction(localAddress string, remoteAddress string, chainId string, frequency int64, expires time.Time, instruction string, strategyId int64, contractAddress string) *MsgUpdateUserInstruction {
	return &MsgUpdateUserInstruction{
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

func (msg *MsgUpdateUserInstruction) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUserInstruction) Type() string {
	return TypeMsgUpdateUserInstruction
}

func (msg *MsgUpdateUserInstruction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.LocalAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUserInstruction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUserInstruction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.LocalAddress)
	if err != nil {
		return sdkerr.Wrapf(sdkerrors.ErrInvalidAddress, "invalid local address (%s)", err)
	}
	return nil
}
