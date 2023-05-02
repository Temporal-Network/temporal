package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"temporal/x/compounder/types"
)

// SetCompoundSettings set a specific compoundSettings in the store from its index
func (k Keeper) SetCompoundSettings(ctx sdk.Context, compoundSettings types.CompoundSettings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingsKeyPrefix))
	b := k.cdc.MustMarshal(&compoundSettings)
	store.Set(types.CompoundSettingsKey(
		compoundSettings.Delegator,
	), b)
}

// GetCompoundSettings returns a compoundSettings from its index
func (k Keeper) GetCompoundSettings(
	ctx sdk.Context,
	delegator string,

) (val types.CompoundSettings, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingsKeyPrefix))

	b := store.Get(types.CompoundSettingsKey(
		delegator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCompoundSettings removes a compoundSettings from the store
func (k Keeper) RemoveCompoundSettings(
	ctx sdk.Context,
	delegator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingsKeyPrefix))
	store.Delete(types.CompoundSettingsKey(
		delegator,
	))
}

// GetAllCompoundSettings returns all compoundSettings
func (k Keeper) GetAllCompoundSettings(ctx sdk.Context) (list []types.CompoundSettings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompoundSettingsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CompoundSettings
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
