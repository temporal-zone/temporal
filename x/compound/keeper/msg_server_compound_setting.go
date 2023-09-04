package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/temporal-zone/temporal/x/compound/types"
	"golang.org/x/exp/slices"
)

func (k msgServer) CreateCompoundSetting(goCtx context.Context, msg *types.MsgCreateCompoundSetting) (*types.MsgCreateCompoundSettingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetCompoundSetting(
		ctx,
		msg.Delegator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "compoundSettings already set, do an update instead")
	}

	err := k.ValidateValidatorSettings(ctx, msg.ValidatorSetting)
	if err != nil {
		return nil, err
	}

	var compoundSetting = types.CompoundSetting{
		Delegator:        msg.Delegator,
		ValidatorSetting: msg.ValidatorSetting,
		AmountToRemain:   msg.AmountToRemain,
		Frequency:        k.CheckFrequency(ctx, msg.Frequency),
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
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "CompoundSettings not found, create them first")
	}

	err := k.ValidateValidatorSettings(ctx, msg.ValidatorSetting)
	if err != nil {
		return nil, err
	}

	// Checks if the msg delegator is the same as the current owner
	if msg.Delegator != valFound.Delegator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var compoundSetting = types.CompoundSetting{
		Delegator:        msg.Delegator,
		ValidatorSetting: msg.ValidatorSetting,
		AmountToRemain:   msg.AmountToRemain,
		Frequency:        k.CheckFrequency(ctx, msg.Frequency),
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
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "CompoundSetting not found")
	}

	// Checks if the msg delegator is the same as the current owner
	if msg.Delegator != valFound.Delegator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCompoundSetting(
		ctx,
		msg.Delegator,
	)

	return &types.MsgDeleteCompoundSettingResponse{}, nil
}

// validateValidatorSettings makes sure ValidatorSetting is valid
func (k msgServer) ValidateValidatorSettings(ctx sdk.Context, validatorSetting []*types.ValidatorSetting) error {
	if validatorSetting == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validatorSetting can not be empty")
	}

	totalPercentToCompound := uint64(0)
	valoperAddresses := make([]string, len(validatorSetting))
	for _, valSetting := range validatorSetting {
		if valSetting.GetPercentToCompound() < 1 || valSetting.GetPercentToCompound() > 100 {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "percentToCompound can not be less than 1 or greater than 100")
		}

		totalPercentToCompound += valSetting.GetPercentToCompound()
		if totalPercentToCompound < 1 || totalPercentToCompound > 100 {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "total percentToCompound across all ValidatorSetting can not be less than 1 or greater than 100")
		}

		valAddress, err := sdk.ValAddressFromBech32(valSetting.ValidatorAddress)
		if err != nil {
			return err
		}

		_, found := k.stakingKeeper.GetValidator(ctx, valAddress)
		if !found {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can not find validator")
		}

		if slices.Contains(valoperAddresses, valSetting.ValidatorAddress) {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator address can not be found in another validator setting")
		}
		valoperAddresses = append(valoperAddresses, valSetting.ValidatorAddress)
	}

	return nil
}
