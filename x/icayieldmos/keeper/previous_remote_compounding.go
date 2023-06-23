package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

// SetPreviousRemoteCompounding set a specific previousRemoteCompounding in the store from its index
func (k Keeper) SetPreviousRemoteCompounding(ctx sdk.Context, previousRemoteCompounding types.PreviousRemoteCompounding) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousRemoteCompoundingKeyPrefix))
	b := k.cdc.MustMarshal(&previousRemoteCompounding)
	store.Set(types.PreviousRemoteCompoundingKey(
		previousRemoteCompounding.RemoteContractCompoundSetting,
	), b)
}

// GetPreviousRemoteCompounding returns a previousRemoteCompounding from its index
func (k Keeper) GetPreviousRemoteCompounding(ctx sdk.Context, remoteContractCompoundSetting uint64) (val types.PreviousRemoteCompounding, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousRemoteCompoundingKeyPrefix))

	b := store.Get(types.PreviousRemoteCompoundingKey(
		remoteContractCompoundSetting,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePreviousRemoteCompounding removes a previousRemoteCompounding from the store
func (k Keeper) RemovePreviousRemoteCompounding(
	ctx sdk.Context,
	remoteContractCompoundSetting uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousRemoteCompoundingKeyPrefix))
	store.Delete(types.PreviousRemoteCompoundingKey(
		remoteContractCompoundSetting,
	))
}

// GetAllPreviousRemoteCompounding returns all previousRemoteCompounding
func (k Keeper) GetAllPreviousRemoteCompounding(ctx sdk.Context) (list []types.PreviousRemoteCompounding) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousRemoteCompoundingKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PreviousRemoteCompounding
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
