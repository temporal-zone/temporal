package types

import (
	sdkerr "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
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

func equalLocalRemoteAddress(localAddress string, remoteAddress string) (bool, error) {
	bech32Prefix, err := deriveBech32Prefix(remoteAddress)
	if err != nil {
		return false, err
	}

	derivedRemoteAddress, err := DeriveRemoteAddress(localAddress, bech32Prefix)
	if err != nil {
		return false, err
	}

	return derivedRemoteAddress == remoteAddress, nil
}

// TODO: This derivation will only work on 44/118 deriv path addresses, will not work on non-44/118 addresses.
func DeriveRemoteAddress(localAddress string, bech32Prefix string) (string, error) {
	_, base64, err := bech32.DecodeAndConvert(localAddress)
	if err != nil {
		return "", err
	}

	remoteAddressConverted, err := bech32.ConvertAndEncode(bech32Prefix, base64)
	if err != nil {
		return "", err
	}

	return remoteAddressConverted, nil
}

// TODO Assumption that all bech32 address have the number 1 in them right after the bech32 prefix, is this true for all?
func deriveBech32Prefix(remoteAddress string) (string, error) {
	index := strings.Index(remoteAddress, "1")
	if index == -1 {
		return "", sdkerr.Wrapf(sdkerrors.ErrInvalidAddress, "not a valid remote address: %s", remoteAddress)
	}

	return remoteAddress[:index], nil
}
