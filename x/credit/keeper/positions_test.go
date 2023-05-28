package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "igmf/testutil/keeper"
	"igmf/testutil/nullify"
	"igmf/x/credit/keeper"
	"igmf/x/credit/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPositions(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Positions {
	items := make([]types.Positions, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetPositions(ctx, items[i])
	}
	return items
}

func TestPositionsGet(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNPositions(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPositions(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPositionsRemove(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNPositions(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePositions(ctx,
			item.Address,
		)
		_, found := keeper.GetPositions(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestPositionsGetAll(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNPositions(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPositions(ctx)),
	)
}
