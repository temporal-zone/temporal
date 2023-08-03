package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"temporal/x/compound/types"
)

// SetCompoundSetting set a specific compoundSetting in the store from its index
func (k Keeper) SetCompoundSetting(ctx sdk.Context, compoundSetting types.CompoundSetting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingKeyPrefix))
	b := k.cdc.MustMarshal(&compoundSetting)
	store.Set(types.CompoundSettingKey(
		compoundSetting.Index123,
	), b)
}

// GetCompoundSetting returns a compoundSetting from its index
func (k Keeper) GetCompoundSetting(
	ctx sdk.Context,
	index123 string,

) (val types.CompoundSetting, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingKeyPrefix))

	b := store.Get(types.CompoundSettingKey(
		index123,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCompoundSetting removes a compoundSetting from the store
func (k Keeper) RemoveCompoundSetting(
	ctx sdk.Context,
	index123 string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingKeyPrefix))
	store.Delete(types.CompoundSettingKey(
		index123,
	))
}

// GetAllCompoundSetting returns all compoundSetting
func (k Keeper) GetAllCompoundSetting(ctx sdk.Context) (list []types.CompoundSetting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CompoundSetting
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
