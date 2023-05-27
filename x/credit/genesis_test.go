package credit_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "igmf/testutil/keeper"
	"igmf/testutil/nullify"
	"igmf/x/credit"
	"igmf/x/credit/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CreditList: []types.Credit{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ModuleInfo: &types.ModuleInfo{
			TotalPositions: 37,
			TotalCredited:  94,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CreditKeeper(t)
	credit.InitGenesis(ctx, *k, genesisState)
	got := credit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CreditList, got.CreditList)
	require.Equal(t, genesisState.ModuleInfo, got.ModuleInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
