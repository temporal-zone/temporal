package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/compound/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.NumberOfCompoundsPerBlock(ctx),
		k.MinimumCompoundFrequency(ctx),
		k.CompoundModuleEnabled(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// NumberOfCompoundsPerBlock returns the NumberOfCompoundsPerBlock param
func (k Keeper) NumberOfCompoundsPerBlock(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyNumberOfCompoundsPerBlock, &res)
	return
}

// MinimumCompoundFrequency returns the MinimumCompoundFrequency param
func (k Keeper) MinimumCompoundFrequency(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyMinimumCompoundFrequency, &res)
	return
}

// CompoundModuleEnabled returns the CompoundModuleEnabled param
func (k Keeper) CompoundModuleEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyCompoundModuleEnabled, &res)
	return
}
