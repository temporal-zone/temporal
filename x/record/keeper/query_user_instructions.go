package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/temporal-zone/temporal/x/record/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserInstructionsAll(goCtx context.Context, req *types.QueryAllUserInstructionsRequest) (*types.QueryAllUserInstructionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userInstructionss []types.UserInstructions
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	userInstructionsStore := prefix.NewStore(store, types.KeyPrefix(types.UserInstructionsKeyPrefix))

	pageRes, err := query.Paginate(userInstructionsStore, req.Pagination, func(key []byte, value []byte) error {
		var userInstructions types.UserInstructions
		if err := k.cdc.Unmarshal(value, &userInstructions); err != nil {
			return err
		}

		userInstructionss = append(userInstructionss, userInstructions)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserInstructionsResponse{UserInstructions: userInstructionss, Pagination: pageRes}, nil
}

func (k Keeper) UserInstructions(goCtx context.Context, req *types.QueryGetUserInstructionsRequest) (*types.QueryGetUserInstructionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetUserInstructions(ctx, req.Address)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUserInstructionsResponse{UserInstructions: val}, nil
}
