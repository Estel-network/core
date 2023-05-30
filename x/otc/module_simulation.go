package otc

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"igmf/testutil/sample"
	otcsimulation "igmf/x/otc/simulation"
	"igmf/x/otc/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = otcsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateTransaction = "op_weight_msg_create_transaction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTransaction int = 100

	opWeightMsgEndTransaction = "op_weight_msg_end_transaction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEndTransaction int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	otcGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&otcGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTransaction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTransaction, &weightMsgCreateTransaction, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTransaction = defaultWeightMsgCreateTransaction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTransaction,
		otcsimulation.SimulateMsgCreateTransaction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEndTransaction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEndTransaction, &weightMsgEndTransaction, nil,
		func(_ *rand.Rand) {
			weightMsgEndTransaction = defaultWeightMsgEndTransaction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEndTransaction,
		otcsimulation.SimulateMsgEndTransaction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
