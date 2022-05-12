package registry

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/testutil/sample"
	registrysimulation "github.com/sonr-io/sonr/x/registry/simulation"
	"github.com/sonr-io/sonr/x/registry/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = registrysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateWhoIs          = "op_weight_msg_who_is"
	defaultWeightMsgCreateWhoIs int = 100

	opWeightMsgUpdateWhoIs          = "op_weight_msg_who_is"
	defaultWeightMsgUpdateWhoIs int = 50

	opWeightMsgDeactivateWhoIs          = "op_weight_msg_who_is"
	defaultWeightMsgDeactivateWhoIs int = 75

	opWeightMsgBuyNameAlias = "op_weight_msg_buy_name_alias"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyNameAlias int = 100

	opWeightMsgBuyAppAlias = "op_weight_msg_buy_app_alias"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyAppAlias int = 100

	opWeightMsgTransferNameAlias = "op_weight_msg_transfer_name_alias"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferNameAlias int = 100

	opWeightMsgTransferAppAlias = "op_weight_msg_transfer_app_alias"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferAppAlias int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	registryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		WhoIsList: []types.WhoIs{
			{
				DidDocument: []byte("did:snr:1"),
				Owner:       sample.AccAddress(),
			},
			{
				DidDocument: []byte("did:snr:2"),
				Owner:       sample.AccAddress(),
			},
		},
		WhoIsCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&registryGenesis)
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

	var weightMsgCreateWhoIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWhoIs, &weightMsgCreateWhoIs, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWhoIs = defaultWeightMsgCreateWhoIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWhoIs,
		registrysimulation.SimulateMsgCreateWhoIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWhoIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWhoIs, &weightMsgUpdateWhoIs, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWhoIs = defaultWeightMsgUpdateWhoIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWhoIs,
		registrysimulation.SimulateMsgUpdateWhoIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeactivateWhoIs int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeactivateWhoIs, &weightMsgDeactivateWhoIs, nil,
		func(_ *rand.Rand) {
			weightMsgDeactivateWhoIs = defaultWeightMsgDeactivateWhoIs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeactivateWhoIs,
		registrysimulation.SimulateMsgDeleteWhoIs(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	var weightMsgBuyNameAlias int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyNameAlias, &weightMsgBuyNameAlias, nil,
		func(_ *rand.Rand) {
			weightMsgBuyNameAlias = defaultWeightMsgBuyNameAlias
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyNameAlias,
		registrysimulation.SimulateMsgBuyNameAlias(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuyAppAlias int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyAppAlias, &weightMsgBuyAppAlias, nil,
		func(_ *rand.Rand) {
			weightMsgBuyAppAlias = defaultWeightMsgBuyAppAlias
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyAppAlias,
		registrysimulation.SimulateMsgBuyAppAlias(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferNameAlias int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferNameAlias, &weightMsgTransferNameAlias, nil,
		func(_ *rand.Rand) {
			weightMsgTransferNameAlias = defaultWeightMsgTransferNameAlias
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferNameAlias,
		registrysimulation.SimulateMsgTransferNameAlias(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferAppAlias int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferAppAlias, &weightMsgTransferAppAlias, nil,
		func(_ *rand.Rand) {
			weightMsgTransferAppAlias = defaultWeightMsgTransferAppAlias
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferAppAlias,
		registrysimulation.SimulateMsgTransferAppAlias(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
