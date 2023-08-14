package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/temporal-zone/temporal/x/compound/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PreviousCompoundAll(goCtx context.Context, req *types.QueryAllPreviousCompoundRequest) (*types.QueryAllPreviousCompoundResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var previousCompounds []types.PreviousCompound
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	previousCompoundStore := prefix.NewStore(store, types.KeyPrefix(types.PreviousCompoundKeyPrefix))

	pageRes, err := query.Paginate(previousCompoundStore, req.Pagination, func(key []byte, value []byte) error {
		var previousCompound types.PreviousCompound
		if err := k.cdc.Unmarshal(value, &previousCompound); err != nil {
			return err
		}

		previousCompounds = append(previousCompounds, previousCompound)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPreviousCompoundResponse{PreviousCompound: previousCompounds, Pagination: pageRes}, nil
}

func (k Keeper) PreviousCompound(goCtx context.Context, req *types.QueryGetPreviousCompoundRequest) (*types.QueryGetPreviousCompoundResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPreviousCompound(
		ctx,
		req.Delegator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPreviousCompoundResponse{PreviousCompound: val}, nil
}
