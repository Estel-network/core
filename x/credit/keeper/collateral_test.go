package keeper_test

import (
	"strconv"
	"testing"

	keepertest "core/testutil/keeper"
	"core/testutil/nullify"
	"core/x/credit/keeper"
	"core/x/credit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCollateral(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Collateral {
	items := make([]types.Collateral, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCollateral(ctx, items[i])
	}
	return items
}

func TestCollateralGet(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNCollateral(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCollateral(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCollateralRemove(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNCollateral(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCollateral(ctx,
			item.Index,
		)
		_, found := keeper.GetCollateral(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCollateralGetAll(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	items := createNCollateral(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCollateral(ctx)),
	)
}
