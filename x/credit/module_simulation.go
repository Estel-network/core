package credit

import (
	"math/rand"

	"core/testutil/sample"
	creditsimulation "core/x/credit/simulation"
	"core/x/credit/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = creditsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeposit int = 100

	opWeightMsgMintCredits = "op_weight_msg_mint_credits"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintCredits int = 100

	opWeightMsgRepayCredits = "op_weight_msg_repay_credits"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayCredits int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	creditGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&creditGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeposit, &weightMsgDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeposit = defaultWeightMsgDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeposit,
		creditsimulation.SimulateMsgDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMintCredits int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintCredits, &weightMsgMintCredits, nil,
		func(_ *rand.Rand) {
			weightMsgMintCredits = defaultWeightMsgMintCredits
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintCredits,
		creditsimulation.SimulateMsgMintCredits(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRepayCredits int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRepayCredits, &weightMsgRepayCredits, nil,
		func(_ *rand.Rand) {
			weightMsgRepayCredits = defaultWeightMsgRepayCredits
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayCredits,
		creditsimulation.SimulateMsgRepayCredits(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeposit,
			defaultWeightMsgDeposit,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				creditsimulation.SimulateMsgDeposit(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMintCredits,
			defaultWeightMsgMintCredits,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				creditsimulation.SimulateMsgMintCredits(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRepayCredits,
			defaultWeightMsgRepayCredits,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				creditsimulation.SimulateMsgRepayCredits(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
