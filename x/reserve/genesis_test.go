package reserve_test

import (
	"testing"

	keepertest "core/testutil/keeper"
	"core/testutil/nullify"
	"core/x/reserve"
	"core/x/reserve/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ReserveKeeper(t)
	reserve.InitGenesis(ctx, *k, genesisState)
	got := reserve.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
