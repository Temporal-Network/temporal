package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"temporal/x/compounder/types"
)

// SetPreviousCompounding set a specific previousCompounding in the store from its index
func (k Keeper) SetPreviousCompounding(ctx sdk.Context, previousCompounding types.PreviousCompounding) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousCompoundingKeyPrefix))
	b := k.cdc.MustMarshal(&previousCompounding)
	store.Set(types.PreviousCompoundingKey(
		previousCompounding.Delegator,
	), b)
}

// GetPreviousCompounding returns a previousCompounding from its index
func (k Keeper) GetPreviousCompounding(ctx sdk.Context, delegator string) (val types.PreviousCompounding, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousCompoundingKeyPrefix))

	b := store.Get(types.PreviousCompoundingKey(
		delegator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePreviousCompounding removes a previousCompounding from the store
func (k Keeper) RemovePreviousCompounding(ctx sdk.Context, delegator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousCompoundingKeyPrefix))
	store.Delete(types.PreviousCompoundingKey(
		delegator,
	))
}

// GetAllPreviousCompounding returns all previousCompounding
func (k Keeper) GetAllPreviousCompounding(ctx sdk.Context) (list []types.PreviousCompounding) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousCompoundingKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PreviousCompounding
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
