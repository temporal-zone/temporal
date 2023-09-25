package keeper

import (
	"context"
	sdkerr "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/record/types"
)

// DeleteUserInstruction will delete UserInstructions based on the fields that are present in MsgDeleteUserInstruction.
// If only the local address is present, it will delete all UserInstructions. If any other fields are populated, it
// will match on their respective fields.
func (k msgServer) DeleteUserInstruction(goCtx context.Context, msg *types.MsgDeleteUserInstruction) (*types.MsgDeleteUserInstructionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetUserInstructions(ctx, msg.LocalAddress)
	if !found {
		return nil, sdkerr.Wrap(sdkerrors.ErrInvalidRequest, "nothing to delete, user instructions not found")
	}

	remoteAddress := msg.GetRemoteAddress()
	chainId := msg.GetChainId()
	contractAddress := msg.GetContractAddress()

	if remoteAddress == "" && chainId == "" && contractAddress == "" {
		k.RemoveUserInstructions(ctx, msg.LocalAddress)
	} else {
		var toKeep []*types.UserInstruction
		for _, userIns := range val.GetUserInstruction() {
			if (remoteAddress == "" || userIns.GetRemoteAddress() == remoteAddress) &&
				(chainId == "" || userIns.GetChainId() == chainId) &&
				(contractAddress == "" || userIns.GetContractAddress() == contractAddress) {
				continue
			}
			toKeep = append(toKeep, userIns)
		}
		val.UserInstruction = toKeep
		k.SetUserInstructions(ctx, val)
	}

	return &types.MsgDeleteUserInstructionResponse{}, nil
}
