package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/record/types"
)

func (k msgServer) UpdateUserInstruction(goCtx context.Context, msg *types.MsgUpdateUserInstruction) (*types.MsgUpdateUserInstructionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateUserInstructionResponse{}, nil
}
