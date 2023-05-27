package credit

import (
	"igmf/x/credit/keeper"
	"igmf/x/credit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined

	k.SetModuleInfo(ctx, genState.ModuleInfo)

	// Set all the credit
	for _, elem := range genState.CreditList {
		k.SetCredit(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all moduleInfo
	moduleInfo, found := k.GetModuleInfo(ctx)
	if found {
		genesis.ModuleInfo = moduleInfo
	}
	genesis.CreditList = k.GetAllCredit(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
