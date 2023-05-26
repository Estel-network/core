package reserve_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "igmf/testutil/keeper"
	"igmf/testutil/nullify"
	"igmf/x/reserve"
	"igmf/x/reserve/types"
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
