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

	if msg.GetRemoteAddress() == "" {
		return sdkerr.Wrapf(sdkerrors.ErrInvalidAddress, "invalid remote address (%s)", msg.GetRemoteAddress())
	}

	addressesEqual, err := equalLocalRemoteAddress(msg.LocalAddress, msg.RemoteAddress)
	if err != nil {
		return err
	}

	if !addressesEqual {
		return sdkerr.Wrapf(sdkerrors.ErrUnauthorized, "local address (%s) and remote address (%s) do not match", msg.GetLocalAddress(), msg.GetRemoteAddress())
	}

	if msg.GetChainId() == "" {
		return sdkerr.Wrapf(sdkerrors.ErrInvalidRequest, "chain id can not be empty (%s)", msg.GetChainId())
	}

	//TODO should this be a module param?
	if msg.GetFrequency() < 3600 {
		return sdkerr.Wrapf(sdkerrors.ErrInvalidRequest, "frequency can not be less than 3600 (%d)", msg.GetFrequency())
	}

	if msg.GetExpires().IsZero() {
		return sdkerr.Wrapf(sdkerrors.ErrInvalidRequest, "expiry can not be zero (%s)", msg.GetExpires().String())
	}

	if msg.GetInstruction() == "" {
		return sdkerr.Wrapf(sdkerrors.ErrInvalidRequest, "instruction can not be empty (%s)", msg.GetInstruction())
	}

	if msg.GetStrategyId() == 0 && msg.GetContractAddress() == "" {
		return sdkerr.Wrapf(sdkerrors.ErrInvalidRequest, "both strategy id and contract address can not be empty")
	}

	return nil
}
