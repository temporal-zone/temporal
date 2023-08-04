package compound

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/temporal-zone/temporal/testutil/sample"
	compoundsimulation "github.com/temporal-zone/temporal/x/compound/simulation"
	"github.com/temporal-zone/temporal/x/compound/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = compoundsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateCompoundSetting = "op_weight_msg_compound_setting"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCompoundSetting int = 100

	opWeightMsgUpdateCompoundSetting = "op_weight_msg_compound_setting"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCompoundSetting int = 100

	opWeightMsgDeleteCompoundSetting = "op_weight_msg_compound_setting"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCompoundSetting int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	compoundGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		CompoundSettingList: []types.CompoundSetting{
			{
				Delegator: sample.AccAddress(),
			},
			{
				Delegator: sample.AccAddress(),
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&compoundGenesis)
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

	var weightMsgCreateCompoundSetting int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCompoundSetting, &weightMsgCreateCompoundSetting, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCompoundSetting = defaultWeightMsgCreateCompoundSetting
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCompoundSetting,
		compoundsimulation.SimulateMsgCreateCompoundSetting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCompoundSetting int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCompoundSetting, &weightMsgUpdateCompoundSetting, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCompoundSetting = defaultWeightMsgUpdateCompoundSetting
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCompoundSetting,
		compoundsimulation.SimulateMsgUpdateCompoundSetting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCompoundSetting int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCompoundSetting, &weightMsgDeleteCompoundSetting, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCompoundSetting = defaultWeightMsgDeleteCompoundSetting
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCompoundSetting,
		compoundsimulation.SimulateMsgDeleteCompoundSetting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateCompoundSetting,
			defaultWeightMsgCreateCompoundSetting,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				compoundsimulation.SimulateMsgCreateCompoundSetting(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateCompoundSetting,
			defaultWeightMsgUpdateCompoundSetting,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				compoundsimulation.SimulateMsgUpdateCompoundSetting(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteCompoundSetting,
			defaultWeightMsgDeleteCompoundSetting,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				compoundsimulation.SimulateMsgDeleteCompoundSetting(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
