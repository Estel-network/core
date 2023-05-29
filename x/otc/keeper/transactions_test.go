package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "igmf/testutil/keeper"
	"igmf/testutil/nullify"
	"igmf/x/otc/keeper"
	"igmf/x/otc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTransactions(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Transactions {
	items := make([]types.Transactions, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTransactions(ctx, items[i])
	}
	return items
}

func TestTransactionsGet(t *testing.T) {
	keeper, ctx := keepertest.OtcKeeper(t)
	items := createNTransactions(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTransactions(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTransactionsRemove(t *testing.T) {
	keeper, ctx := keepertest.OtcKeeper(t)
	items := createNTransactions(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTransactions(ctx,
			item.Index,
		)
		_, found := keeper.GetTransactions(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTransactionsGetAll(t *testing.T) {
	keeper, ctx := keepertest.OtcKeeper(t)
	items := createNTransactions(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTransactions(ctx)),
	)
}
