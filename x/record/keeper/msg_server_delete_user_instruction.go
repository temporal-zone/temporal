package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/record/types"
)

func (k msgServer) DeleteUserInstruction(goCtx context.Context, msg *types.MsgDeleteUserInstruction) (*types.MsgDeleteUserInstructionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteUserInstructionResponse{}, nil
}
