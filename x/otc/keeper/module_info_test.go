package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "igmf/testutil/keeper"
	"igmf/testutil/nullify"
	"igmf/x/otc/keeper"
	"igmf/x/otc/types"
)

func createTestModuleInfo(keeper *keeper.Keeper, ctx sdk.Context) types.ModuleInfo {
	item := types.ModuleInfo{}
	keeper.SetModuleInfo(ctx, item)
	return item
}

func TestModuleInfoGet(t *testing.T) {
	keeper, ctx := keepertest.OtcKeeper(t)
	item := createTestModuleInfo(keeper, ctx)
	rst, found := keeper.GetModuleInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestModuleInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.OtcKeeper(t)
	createTestModuleInfo(keeper, ctx)
	keeper.RemoveModuleInfo(ctx)
	_, found := keeper.GetModuleInfo(ctx)
	require.False(t, found)
}
