package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/record/types"
)

// SetUserInstructions set a specific userInstructions in the store from its index
func (k Keeper) SetUserInstructions(ctx sdk.Context, userInstructions types.UserInstructions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInstructionsKeyPrefix))
	b := k.cdc.MustMarshal(&userInstructions)
	store.Set(types.UserInstructionsKey(userInstructions.Address), b)
}

// GetUserInstructions returns a userInstructions from its index
func (k Keeper) GetUserInstructions(ctx sdk.Context, address string) (val types.UserInstructions, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInstructionsKeyPrefix))

	b := store.Get(types.UserInstructionsKey(address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserInstructions removes a userInstructions from the store
func (k Keeper) RemoveUserInstructions(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInstructionsKeyPrefix))
	store.Delete(types.UserInstructionsKey(address))
}

// GetAllUserInstructions returns all userInstructions
func (k Keeper) GetAllUserInstructions(ctx sdk.Context) (list []types.UserInstructions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserInstructionsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserInstructions
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
