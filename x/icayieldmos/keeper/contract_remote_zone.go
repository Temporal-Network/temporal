package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/x/icayieldmos/types"
)

// GetContractRemoteZoneCount get the total number of contractRemoteZone
func (k Keeper) GetContractRemoteZoneCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ContractRemoteZoneCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetContractRemoteZoneCount set the total number of contractRemoteZone
func (k Keeper) SetContractRemoteZoneCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ContractRemoteZoneCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendContractRemoteZone appends a contractRemoteZone in the store with a new id and update the count
func (k Keeper) AppendContractRemoteZone(
	ctx sdk.Context,
	contractRemoteZone types.ContractRemoteZone,
) uint64 {
	// Create the contractRemoteZone
	count := k.GetContractRemoteZoneCount(ctx)

	// Set the ID of the appended value
	contractRemoteZone.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractRemoteZoneKey))
	appendedValue := k.cdc.MustMarshal(&contractRemoteZone)
	store.Set(GetContractRemoteZoneIDBytes(contractRemoteZone.Id), appendedValue)

	// Update contractRemoteZone count
	k.SetContractRemoteZoneCount(ctx, count+1)

	return count
}

// SetContractRemoteZone set a specific contractRemoteZone in the store
func (k Keeper) SetContractRemoteZone(ctx sdk.Context, contractRemoteZone types.ContractRemoteZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractRemoteZoneKey))
	b := k.cdc.MustMarshal(&contractRemoteZone)
	store.Set(GetContractRemoteZoneIDBytes(contractRemoteZone.Id), b)
}

// GetContractRemoteZone returns a contractRemoteZone from its id
func (k Keeper) GetContractRemoteZone(ctx sdk.Context, id uint64) (val types.ContractRemoteZone, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractRemoteZoneKey))
	b := store.Get(GetContractRemoteZoneIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContractRemoteZone removes a contractRemoteZone from the store
func (k Keeper) RemoveContractRemoteZone(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractRemoteZoneKey))
	store.Delete(GetContractRemoteZoneIDBytes(id))
}

// GetAllContractRemoteZone returns all contractRemoteZone
func (k Keeper) GetAllContractRemoteZone(ctx sdk.Context) (list []types.ContractRemoteZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractRemoteZoneKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ContractRemoteZone
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetContractRemoteZoneIDBytes returns the byte representation of the ID
func GetContractRemoteZoneIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetContractRemoteZoneIDFromBytes returns ID in uint64 format from a byte array
func GetContractRemoteZoneIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
