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
	store.Set(types.CompoundSettingKey(compoundSetting.Delegator), b)
}

// GetCompoundSetting returns a compoundSetting from its index
func (k Keeper) GetCompoundSetting(ctx sdk.Context, delegator string) (val types.CompoundSetting, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingKeyPrefix))

	b := store.Get(types.CompoundSettingKey(delegator))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCompoundSetting removes a compoundSetting from the store
func (k Keeper) RemoveCompoundSetting(ctx sdk.Context, delegator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingKeyPrefix))
	store.Delete(types.CompoundSettingKey(delegator))
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
