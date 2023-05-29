package otc_test

import (
	"testing"

	keepertest "igmf/testutil/keeper"
	"igmf/testutil/nullify"
	"igmf/x/otc"
	"igmf/x/otc/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TransactionsList: []types.Transactions{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ModuleInfo: types.ModuleInfo{
			ModuleIndex: 97,
			ServiceFee:  91,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OtcKeeper(t)
	otc.InitGenesis(ctx, *k, genesisState)
	got := otc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TransactionsList, got.TransactionsList)
	require.Equal(t, genesisState.ModuleInfo, got.ModuleInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
