package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCompoundSetting = "create_compound_setting"
	TypeMsgUpdateCompoundSetting = "update_compound_setting"
	TypeMsgDeleteCompoundSetting = "delete_compound_setting"
)

var _ sdk.Msg = &MsgCreateCompoundSetting{}

func NewMsgCreateCompoundSetting(
	delegator string,
	validatorSetting []*ValidatorSetting,
	amountToRemain sdk.Coin,
	frequency uint64,

) *MsgCreateCompoundSetting {
	return &MsgCreateCompoundSetting{
		Delegator:        delegator,
		ValidatorSetting: validatorSetting,
		AmountToRemain:   amountToRemain,
		Frequency:        frequency,
	}
}

func (msg *MsgCreateCompoundSetting) Route() string {
	return RouterKey
}

func (msg *MsgCreateCompoundSetting) Type() string {
	return TypeMsgCreateCompoundSetting
}

func (msg *MsgCreateCompoundSetting) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

func (msg *MsgCreateCompoundSetting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCompoundSetting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCompoundSetting{}

func NewMsgUpdateCompoundSetting(
	delegator string,
	validatorSetting []*ValidatorSetting,
	amountToRemain sdk.Coin,
	frequency uint64,

) *MsgUpdateCompoundSetting {
	return &MsgUpdateCompoundSetting{
		Delegator:        delegator,
		ValidatorSetting: validatorSetting,
		AmountToRemain:   amountToRemain,
		Frequency:        frequency,
	}
}

func (msg *MsgUpdateCompoundSetting) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCompoundSetting) Type() string {
	return TypeMsgUpdateCompoundSetting
}

func (msg *MsgUpdateCompoundSetting) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

func (msg *MsgUpdateCompoundSetting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCompoundSetting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCompoundSetting{}

func NewMsgDeleteCompoundSetting(
	delegator string,

) *MsgDeleteCompoundSetting {
	return &MsgDeleteCompoundSetting{
		Delegator: delegator,
	}
}
func (msg *MsgDeleteCompoundSetting) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCompoundSetting) Type() string {
	return TypeMsgDeleteCompoundSetting
}

func (msg *MsgDeleteCompoundSetting) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

func (msg *MsgDeleteCompoundSetting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCompoundSetting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}
	return nil
}
