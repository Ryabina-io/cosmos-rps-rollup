package rps

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"rps/testutil/sample"
	rpssimulation "rps/x/rps/simulation"
	"rps/x/rps/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = rpssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateGame = "op_weight_msg_create_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGame int = 100

	opWeightMsgJoinGame = "op_weight_msg_join_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgJoinGame int = 100

	opWeightMsgRevealGame = "op_weight_msg_reveal_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevealGame int = 100

	opWeightMsgRemoveGame = "op_weight_msg_remove_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveGame int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rpsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rpsGenesis)
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

	var weightMsgCreateGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateGame, &weightMsgCreateGame, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGame = defaultWeightMsgCreateGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGame,
		rpssimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgJoinGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgJoinGame, &weightMsgJoinGame, nil,
		func(_ *rand.Rand) {
			weightMsgJoinGame = defaultWeightMsgJoinGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgJoinGame,
		rpssimulation.SimulateMsgJoinGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevealGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevealGame, &weightMsgRevealGame, nil,
		func(_ *rand.Rand) {
			weightMsgRevealGame = defaultWeightMsgRevealGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevealGame,
		rpssimulation.SimulateMsgRevealGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveGame, &weightMsgRemoveGame, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveGame = defaultWeightMsgRemoveGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveGame,
		rpssimulation.SimulateMsgRemoveGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
