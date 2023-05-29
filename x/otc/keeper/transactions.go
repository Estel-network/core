package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"igmf/x/otc/types"
)

// SetTransactions set a specific transactions in the store from its index
func (k Keeper) SetTransactions(ctx sdk.Context, transactions types.Transactions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionsKeyPrefix))
	b := k.cdc.MustMarshal(&transactions)
	store.Set(types.TransactionsKey(
		transactions.Index,
	), b)
}

// GetTransactions returns a transactions from its index
func (k Keeper) GetTransactions(
	ctx sdk.Context,
	index string,

) (val types.Transactions, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionsKeyPrefix))

	b := store.Get(types.TransactionsKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTransactions removes a transactions from the store
func (k Keeper) RemoveTransactions(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionsKeyPrefix))
	store.Delete(types.TransactionsKey(
		index,
	))
}

// GetAllTransactions returns all transactions
func (k Keeper) GetAllTransactions(ctx sdk.Context) (list []types.Transactions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Transactions
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
