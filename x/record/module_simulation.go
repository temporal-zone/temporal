package record

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/temporal-zone/temporal/testutil/sample"
	recordsimulation "github.com/temporal-zone/temporal/x/record/simulation"
	"github.com/temporal-zone/temporal/x/record/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = recordsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateUserInstruction = "op_weight_msg_create_user_instruction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUserInstruction int = 100

	opWeightMsgDeleteUserInstruction = "op_weight_msg_delete_user_instruction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUserInstruction int = 100

	opWeightMsgUpdateUserInstruction = "op_weight_msg_update_user_instruction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUserInstruction int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	recordGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&recordGenesis)
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

	var weightMsgCreateUserInstruction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateUserInstruction, &weightMsgCreateUserInstruction, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUserInstruction = defaultWeightMsgCreateUserInstruction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUserInstruction,
		recordsimulation.SimulateMsgCreateUserInstruction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUserInstruction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteUserInstruction, &weightMsgDeleteUserInstruction, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUserInstruction = defaultWeightMsgDeleteUserInstruction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUserInstruction,
		recordsimulation.SimulateMsgDeleteUserInstruction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUserInstruction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateUserInstruction, &weightMsgUpdateUserInstruction, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUserInstruction = defaultWeightMsgUpdateUserInstruction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUserInstruction,
		recordsimulation.SimulateMsgUpdateUserInstruction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateUserInstruction,
			defaultWeightMsgCreateUserInstruction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				recordsimulation.SimulateMsgCreateUserInstruction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteUserInstruction,
			defaultWeightMsgDeleteUserInstruction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				recordsimulation.SimulateMsgDeleteUserInstruction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateUserInstruction,
			defaultWeightMsgUpdateUserInstruction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				recordsimulation.SimulateMsgUpdateUserInstruction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
