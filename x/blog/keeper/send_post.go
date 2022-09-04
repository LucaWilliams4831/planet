package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/blog/types"
)

// GetSendPostCount get the total number of sendPost
func (k Keeper) GetSendPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SendPostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSendPostCount set the total number of sendPost
func (k Keeper) SetSendPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SendPostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSendPost appends a sendPost in the store with a new id and update the count
func (k Keeper) AppendSendPost(
	ctx sdk.Context,
	sendPost types.SendPost,
) uint64 {
	// Create the sendPost
	count := k.GetSendPostCount(ctx)

	// Set the ID of the appended value
	sendPost.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendPostKey))
	appendedValue := k.cdc.MustMarshal(&sendPost)
	store.Set(GetSendPostIDBytes(sendPost.Id), appendedValue)

	// Update sendPost count
	k.SetSendPostCount(ctx, count+1)

	return count
}

// SetSendPost set a specific sendPost in the store
func (k Keeper) SetSendPost(ctx sdk.Context, sendPost types.SendPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendPostKey))
	b := k.cdc.MustMarshal(&sendPost)
	store.Set(GetSendPostIDBytes(sendPost.Id), b)
}

// GetSendPost returns a sendPost from its id
func (k Keeper) GetSendPost(ctx sdk.Context, id uint64) (val types.SendPost, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendPostKey))
	b := store.Get(GetSendPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSendPost removes a sendPost from the store
func (k Keeper) RemoveSendPost(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendPostKey))
	store.Delete(GetSendPostIDBytes(id))
}

// GetAllSendPost returns all sendPost
func (k Keeper) GetAllSendPost(ctx sdk.Context) (list []types.SendPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendPostKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SendPost
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSendPostIDBytes returns the byte representation of the ID
func GetSendPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSendPostIDFromBytes returns ID in uint64 format from a byte array
func GetSendPostIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
