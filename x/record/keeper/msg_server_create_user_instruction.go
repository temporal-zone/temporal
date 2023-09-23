package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/record/types"
)

func (k msgServer) CreateUserInstruction(goCtx context.Context, msg *types.MsgCreateUserInstruction) (*types.MsgCreateUserInstructionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateUserInstructionResponse{}, nil
}
