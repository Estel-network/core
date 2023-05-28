package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"igmf/x/credit/types"
)

// SetPositions set a specific positions in the store from its index
func (k Keeper) SetPositions(ctx sdk.Context, positions types.Positions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionsKeyPrefix))
	b := k.cdc.MustMarshal(&positions)
	store.Set(types.PositionsKey(
		positions.Address,
	), b)
}

// GetPositions returns a positions from its index
func (k Keeper) GetPositions(
	ctx sdk.Context,
	address string,

) (val types.Positions, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionsKeyPrefix))

	b := store.Get(types.PositionsKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePositions removes a positions from the store
func (k Keeper) RemovePositions(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionsKeyPrefix))
	store.Delete(types.PositionsKey(
		address,
	))
}

// GetAllPositions returns all positions
func (k Keeper) GetAllPositions(ctx sdk.Context) (list []types.Positions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Positions
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
