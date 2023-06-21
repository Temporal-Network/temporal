package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Temporal-Network/temporal/x/icayieldmos/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// GetRemoteContractCompoundSettingsCount get the total number of remoteContractCompoundSettings
func (k Keeper) GetRemoteContractCompoundSettingsCount(ctx sdk.Context) uint64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RemoteContractCompoundSettingsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRemoteContractCompoundSettingsCount set the total number of remoteContractCompoundSettings
func (k Keeper) SetRemoteContractCompoundSettingsCount(ctx sdk.Context, count uint64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RemoteContractCompoundSettingsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRemoteContractCompoundSettings appends a remoteContractCompoundSettings in the store with a new id and update the count
func (k Keeper) AppendRemoteContractCompoundSettings(
    ctx sdk.Context,
    remoteContractCompoundSettings types.RemoteContractCompoundSettings,
) uint64 {
	// Create the remoteContractCompoundSettings
    count := k.GetRemoteContractCompoundSettingsCount(ctx)

    // Set the ID of the appended value
    remoteContractCompoundSettings.Id = count

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RemoteContractCompoundSettingsKey))
    appendedValue := k.cdc.MustMarshal(&remoteContractCompoundSettings)
    store.Set(GetRemoteContractCompoundSettingsIDBytes(remoteContractCompoundSettings.Id), appendedValue)

    // Update remoteContractCompoundSettings count
    k.SetRemoteContractCompoundSettingsCount(ctx, count+1)

    return count
}

// SetRemoteContractCompoundSettings set a specific remoteContractCompoundSettings in the store
func (k Keeper) SetRemoteContractCompoundSettings(ctx sdk.Context, remoteContractCompoundSettings types.RemoteContractCompoundSettings) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RemoteContractCompoundSettingsKey))
	b := k.cdc.MustMarshal(&remoteContractCompoundSettings)
	store.Set(GetRemoteContractCompoundSettingsIDBytes(remoteContractCompoundSettings.Id), b)
}

// GetRemoteContractCompoundSettings returns a remoteContractCompoundSettings from its id
func (k Keeper) GetRemoteContractCompoundSettings(ctx sdk.Context, id uint64) (val types.RemoteContractCompoundSettings, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RemoteContractCompoundSettingsKey))
	b := store.Get(GetRemoteContractCompoundSettingsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRemoteContractCompoundSettings removes a remoteContractCompoundSettings from the store
func (k Keeper) RemoveRemoteContractCompoundSettings(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RemoteContractCompoundSettingsKey))
	store.Delete(GetRemoteContractCompoundSettingsIDBytes(id))
}

// GetAllRemoteContractCompoundSettings returns all remoteContractCompoundSettings
func (k Keeper) GetAllRemoteContractCompoundSettings(ctx sdk.Context) (list []types.RemoteContractCompoundSettings) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RemoteContractCompoundSettingsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RemoteContractCompoundSettings
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

// GetRemoteContractCompoundSettingsIDBytes returns the byte representation of the ID
func GetRemoteContractCompoundSettingsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRemoteContractCompoundSettingsIDFromBytes returns ID in uint64 format from a byte array
func GetRemoteContractCompoundSettingsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
