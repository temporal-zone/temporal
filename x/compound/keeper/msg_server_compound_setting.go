package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/temporal-zone/temporal/x/compound/types"
)

func (k msgServer) CreateCompoundSetting(goCtx context.Context, msg *types.MsgCreateCompoundSetting) (*types.MsgCreateCompoundSettingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetCompoundSetting(
		ctx,
		msg.Delegator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var compoundSetting = types.CompoundSetting{
		Delegator:        msg.Delegator,
		ValidatorSetting: msg.ValidatorSetting,
		AmountToRemain:   msg.AmountToRemain,
		Frequency:        CheckFrequency(msg.Frequency),
	}

	k.SetCompoundSetting(
		ctx,
		compoundSetting,
	)
	return &types.MsgCreateCompoundSettingResponse{}, nil
}

func (k msgServer) UpdateCompoundSetting(goCtx context.Context, msg *types.MsgUpdateCompoundSetting) (*types.MsgUpdateCompoundSettingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCompoundSetting(
		ctx,
		msg.Delegator,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg delegator is the same as the current owner
	if msg.Delegator != valFound.Delegator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var compoundSetting = types.CompoundSetting{
		Delegator:        msg.Delegator,
		ValidatorSetting: msg.ValidatorSetting,
		AmountToRemain:   msg.AmountToRemain,
		Frequency:        CheckFrequency(msg.Frequency),
	}

	k.SetCompoundSetting(ctx, compoundSetting)

	return &types.MsgUpdateCompoundSettingResponse{}, nil
}

func (k msgServer) DeleteCompoundSetting(goCtx context.Context, msg *types.MsgDeleteCompoundSetting) (*types.MsgDeleteCompoundSettingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCompoundSetting(
		ctx,
		msg.Delegator,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg delegator is the same as the current owner
	if msg.Delegator != valFound.Delegator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCompoundSetting(
		ctx,
		msg.Delegator,
	)

	return &types.MsgDeleteCompoundSettingResponse{}, nil
}

// TODO: CheckFrequency needs test coverage
// CheckFrequency checks to make sure frequency should be no less than X seconds.
func CheckFrequency(onceEvery uint64) uint64 {
	// TODO: Change minimumCompoundFrequency to be a module param
	minimumCompoundFrequency := uint64(600)
	if onceEvery < minimumCompoundFrequency {
		return minimumCompoundFrequency
	}

	return onceEvery
}
