package otc_test

import (
	"testing"

	keepertest "core/testutil/keeper"
	"core/testutil/nullify"
	"core/x/otc"
	"core/x/otc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OtcKeeper(t)
	otc.InitGenesis(ctx, *k, genesisState)
	got := otc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
