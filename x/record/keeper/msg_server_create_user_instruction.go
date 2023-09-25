package keeper

import (
	"context"
	sdkerr "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/temporal-zone/temporal/x/record/types"
)

func (k msgServer) CreateUserInstruction(goCtx context.Context, msg *types.MsgCreateUserInstruction) (*types.MsgCreateUserInstructionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Expires.Before(ctx.BlockTime()) {
		return nil, sdkerr.Wrap(sdkerrors.ErrInvalidRequest, "expiry is before now")
	}

	userInstruction := types.UserInstruction{
		LocalAddress:    msg.GetLocalAddress(),
		RemoteAddress:   msg.GetRemoteAddress(),
		ChainId:         msg.GetChainId(),
		Frequency:       msg.GetFrequency(),
		Created:         ctx.BlockTime().UTC(),
		Expires:         msg.GetExpires(),
		Instruction:     msg.GetInstruction(),
		StrategyId:      msg.GetStrategyId(),
		ContractAddress: msg.GetContractAddress(),
	}

	val, found := k.GetUserInstructions(ctx, msg.LocalAddress)
	if !found {
		val.Address = userInstruction.LocalAddress
		val.UserInstruction = make([]*types.UserInstruction, 0, 1)
	} else {
		for _, userIns := range val.GetUserInstruction() {
			if equalUserInstruction(userIns, &userInstruction) {
				return nil, sdkerr.Wrap(sdkerrors.ErrInvalidRequest, "duplicate user instruction found, do an update instead")
			}
		}
	}

	val.UserInstruction = append(val.GetUserInstruction(), &userInstruction)

	k.SetUserInstructions(ctx, val)

	return &types.MsgCreateUserInstructionResponse{}, nil
}

// A helper function to compare user instructions.
func equalUserInstruction(ins1 *types.UserInstruction, ins2 *types.UserInstruction) bool {
	return ins1.LocalAddress == ins2.LocalAddress &&
		ins1.RemoteAddress == ins2.RemoteAddress &&
		ins1.ChainId == ins2.ChainId &&
		ins1.StrategyId == ins2.StrategyId &&
		ins1.ContractAddress == ins2.ContractAddress
}
