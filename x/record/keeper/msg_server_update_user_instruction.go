package keeper

import (
	"context"
	sdkerr "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/record/types"
)

func (k msgServer) UpdateUserInstruction(goCtx context.Context, msg *types.MsgUpdateUserInstruction) (*types.MsgUpdateUserInstructionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Expires.Before(ctx.BlockTime()) {
		return nil, sdkerr.Wrap(sdkerrors.ErrInvalidRequest, "expiry is before now")
	}

	val, found := k.GetUserInstructions(ctx, msg.LocalAddress)
	if !found {
		return nil, sdkerr.Wrap(sdkerrors.ErrInvalidRequest, "nothing to update, user instructions not found")
	}

	userInstruction := types.UserInstruction{
		LocalAddress:    msg.GetLocalAddress(),
		RemoteAddress:   msg.GetRemoteAddress(),
		ChainId:         msg.GetChainId(),
		Frequency:       msg.GetFrequency(),
		Created:         time.Time{},
		Expires:         msg.GetExpires(),
		Instruction:     msg.GetInstruction(),
		StrategyId:      msg.GetStrategyId(),
		ContractAddress: msg.GetContractAddress(),
	}

	for i, userIns := range val.GetUserInstruction() {
		if equalUserInstruction(userIns, &userInstruction) {

			userInstruction.Created = userIns.Created
			val.GetUserInstruction()[i] = &userInstruction

			k.SetUserInstructions(ctx, val)

			break
		}

		return nil, sdkerr.Wrap(sdkerrors.ErrInvalidRequest, "nothing to update, user instruction not found")
	}

	return &types.MsgUpdateUserInstructionResponse{}, nil
}
