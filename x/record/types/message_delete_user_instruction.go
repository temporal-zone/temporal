package types

import (
	sdkerr "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteUserInstruction = "delete_user_instruction"

var _ sdk.Msg = &MsgDeleteUserInstruction{}

func NewMsgDeleteUserInstruction(localAddress string, remoteAddress string, chainId string, contractAddress string) *MsgDeleteUserInstruction {
	return &MsgDeleteUserInstruction{
		LocalAddress:    localAddress,
		RemoteAddress:   remoteAddress,
		ChainId:         chainId,
		ContractAddress: contractAddress,
	}
}

func (msg *MsgDeleteUserInstruction) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUserInstruction) Type() string {
	return TypeMsgDeleteUserInstruction
}

func (msg *MsgDeleteUserInstruction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.LocalAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUserInstruction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUserInstruction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.LocalAddress)
	if err != nil {
		return sdkerr.Wrapf(sdkerrors.ErrInvalidAddress, "invalid local address (%s)", err)
	}
	return nil
}
