package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "core/testutil/keeper"
	"core/testutil/nullify"
	"core/x/credit/keeper"
	"core/x/credit/types"
)

func createTestModuleInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ModuleInfo {
	item := types.ModuleInfo{}
	keeper.SetModuleInfo(ctx, item)
	return item
}

func TestModuleInfoGet(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	item := createTestModuleInfo(keeper, ctx)
	rst, found := keeper.GetModuleInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestModuleInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	createTestModuleInfo(keeper, ctx)
	keeper.RemoveModuleInfo(ctx)
	_, found := keeper.GetModuleInfo(ctx)
	require.False(t, found)
}
