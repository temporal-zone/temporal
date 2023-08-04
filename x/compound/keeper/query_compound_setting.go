package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"temporal/x/compound/types"
)

func (k Keeper) CompoundSettingAll(goCtx context.Context, req *types.QueryAllCompoundSettingRequest) (*types.QueryAllCompoundSettingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var compoundSettings []types.CompoundSetting
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	compoundSettingStore := prefix.NewStore(store, types.KeyPrefix(types.CompoundSettingKeyPrefix))

	pageRes, err := query.Paginate(compoundSettingStore, req.Pagination, func(key []byte, value []byte) error {
		var compoundSetting types.CompoundSetting
		if err := k.cdc.Unmarshal(value, &compoundSetting); err != nil {
			return err
		}

		compoundSettings = append(compoundSettings, compoundSetting)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCompoundSettingResponse{CompoundSetting: compoundSettings, Pagination: pageRes}, nil
}

func (k Keeper) CompoundSetting(goCtx context.Context, req *types.QueryGetCompoundSettingRequest) (*types.QueryGetCompoundSettingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetCompoundSetting(
		ctx,
		req.Delegator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCompoundSettingResponse{CompoundSetting: val}, nil
}
