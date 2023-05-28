package credit_test

import (
	"testing"

	keepertest "igmf/testutil/keeper"
	"igmf/testutil/nullify"
	"igmf/x/credit"
	"igmf/x/credit/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ModuleInfo: types.ModuleInfo{
			TotalPositions: 25,
			TotalCredited:  91,
		},
		CreditList: []types.Credit{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		PositionsList: []types.Positions{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CreditKeeper(t)
	credit.InitGenesis(ctx, *k, genesisState)
	got := credit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.ModuleInfo, got.ModuleInfo)
	require.ElementsMatch(t, genesisState.CreditList, got.CreditList)
	require.ElementsMatch(t, genesisState.PositionsList, got.PositionsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
