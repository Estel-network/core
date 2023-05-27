package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"igmf/x/credit/types"
)

// SetModuleInfo set moduleInfo in the store
func (k Keeper) SetModuleInfo(ctx sdk.Context, moduleInfo types.ModuleInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModuleInfoKey))
	b := k.cdc.MustMarshal(&moduleInfo)
	store.Set([]byte{0}, b)
}

// GetModuleInfo returns moduleInfo
func (k Keeper) GetModuleInfo(ctx sdk.Context) (val types.ModuleInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModuleInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveModuleInfo removes moduleInfo from the store
func (k Keeper) RemoveModuleInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModuleInfoKey))
	store.Delete([]byte{0})
}
