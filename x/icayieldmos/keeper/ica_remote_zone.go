package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// GetICARemoteZoneCount get the total number of iCARemoteZone
func (k Keeper) GetICARemoteZoneCount(ctx sdk.Context) uint64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ICARemoteZoneCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetICARemoteZoneCount set the total number of iCARemoteZone
func (k Keeper) SetICARemoteZoneCount(ctx sdk.Context, count uint64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ICARemoteZoneCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendICARemoteZone appends a iCARemoteZone in the store with a new id and update the count
func (k Keeper) AppendICARemoteZone(
    ctx sdk.Context,
    iCARemoteZone types.ICARemoteZone,
) uint64 {
	// Create the iCARemoteZone
    count := k.GetICARemoteZoneCount(ctx)

    // Set the ID of the appended value
    iCARemoteZone.Id = count

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ICARemoteZoneKey))
    appendedValue := k.cdc.MustMarshal(&iCARemoteZone)
    store.Set(GetICARemoteZoneIDBytes(iCARemoteZone.Id), appendedValue)

    // Update iCARemoteZone count
    k.SetICARemoteZoneCount(ctx, count+1)

    return count
}

// SetICARemoteZone set a specific iCARemoteZone in the store
func (k Keeper) SetICARemoteZone(ctx sdk.Context, iCARemoteZone types.ICARemoteZone) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ICARemoteZoneKey))
	b := k.cdc.MustMarshal(&iCARemoteZone)
	store.Set(GetICARemoteZoneIDBytes(iCARemoteZone.Id), b)
}

// GetICARemoteZone returns a iCARemoteZone from its id
func (k Keeper) GetICARemoteZone(ctx sdk.Context, id uint64) (val types.ICARemoteZone, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ICARemoteZoneKey))
	b := store.Get(GetICARemoteZoneIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveICARemoteZone removes a iCARemoteZone from the store
func (k Keeper) RemoveICARemoteZone(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ICARemoteZoneKey))
	store.Delete(GetICARemoteZoneIDBytes(id))
}

// GetAllICARemoteZone returns all iCARemoteZone
func (k Keeper) GetAllICARemoteZone(ctx sdk.Context) (list []types.ICARemoteZone) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ICARemoteZoneKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ICARemoteZone
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

// GetICARemoteZoneIDBytes returns the byte representation of the ID
func GetICARemoteZoneIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetICARemoteZoneIDFromBytes returns ID in uint64 format from a byte array
func GetICARemoteZoneIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
