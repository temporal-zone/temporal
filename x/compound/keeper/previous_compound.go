package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/compound/types"
)

// SetPreviousCompound set a specific previousCompound in the store from its index
func (k Keeper) SetPreviousCompound(ctx sdk.Context, previousCompound types.PreviousCompound) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousCompoundKeyPrefix))
	b := k.cdc.MustMarshal(&previousCompound)
	store.Set(types.PreviousCompoundKey(
		previousCompound.Delegator,
	), b)
}

// GetPreviousCompound returns a previousCompound from its index
func (k Keeper) GetPreviousCompound(ctx sdk.Context, delegator string) (val types.PreviousCompound, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousCompoundKeyPrefix))

	b := store.Get(types.PreviousCompoundKey(
		delegator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePreviousCompound removes a previousCompound from the store
func (k Keeper) RemovePreviousCompound(ctx sdk.Context, delegator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousCompoundKeyPrefix))
	store.Delete(types.PreviousCompoundKey(delegator))
}

// GetAllPreviousCompound returns all previousCompound
func (k Keeper) GetAllPreviousCompound(ctx sdk.Context) (list []types.PreviousCompound) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousCompoundKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PreviousCompound
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
