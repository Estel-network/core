package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "igmf/testutil/keeper"
	"igmf/x/reserve/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ReserveKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}