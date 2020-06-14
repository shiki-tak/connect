package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the ibc store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// NewKeeper creates a ibc keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) Get(ctx sdk.Context, key string) (string, error) {
	// store := ctx.KVStore(k.storeKey)
	// var item /* TODO: Fill out this type */
	// byteKey := []byte(key)
	// err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	// if err != nil {
	// 	return nil, err
	// }
	return "", nil
}

func (k Keeper) set(ctx sdk.Context, key string, value string /* TODO: fill out this type */) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
	store.Set([]byte(key), bz)
}

func (k Keeper) delete(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}
