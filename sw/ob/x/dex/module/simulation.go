package dex

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"ob/testutil/sample"
	dexsimulation "ob/x/dex/simulation"
	"ob/x/dex/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dexGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MarketMap: []types.Market{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, OrderMap: []types.Order{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dexGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateMarket          = "op_weight_msg_dex"
		defaultWeightMsgCreateMarket int = 100
	)

	var weightMsgCreateMarket int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMarket, &weightMsgCreateMarket, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMarket = defaultWeightMsgCreateMarket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMarket,
		dexsimulation.SimulateMsgCreateMarket(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateMarket          = "op_weight_msg_dex"
		defaultWeightMsgUpdateMarket int = 100
	)

	var weightMsgUpdateMarket int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMarket, &weightMsgUpdateMarket, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMarket = defaultWeightMsgUpdateMarket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMarket,
		dexsimulation.SimulateMsgUpdateMarket(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteMarket          = "op_weight_msg_dex"
		defaultWeightMsgDeleteMarket int = 100
	)

	var weightMsgDeleteMarket int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteMarket, &weightMsgDeleteMarket, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMarket = defaultWeightMsgDeleteMarket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMarket,
		dexsimulation.SimulateMsgDeleteMarket(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateOrder          = "op_weight_msg_dex"
		defaultWeightMsgCreateOrder int = 100
	)

	var weightMsgCreateOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateOrder, &weightMsgCreateOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOrder = defaultWeightMsgCreateOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOrder,
		dexsimulation.SimulateMsgCreateOrder(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateOrder          = "op_weight_msg_dex"
		defaultWeightMsgUpdateOrder int = 100
	)

	var weightMsgUpdateOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateOrder, &weightMsgUpdateOrder, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOrder = defaultWeightMsgUpdateOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOrder,
		dexsimulation.SimulateMsgUpdateOrder(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteOrder          = "op_weight_msg_dex"
		defaultWeightMsgDeleteOrder int = 100
	)

	var weightMsgDeleteOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteOrder, &weightMsgDeleteOrder, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteOrder = defaultWeightMsgDeleteOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteOrder,
		dexsimulation.SimulateMsgDeleteOrder(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgRegisterPairs          = "op_weight_msg_dex"
		defaultWeightMsgRegisterPairs int = 100
	)

	var weightMsgRegisterPairs int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterPairs, &weightMsgRegisterPairs, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterPairs = defaultWeightMsgRegisterPairs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterPairs,
		dexsimulation.SimulateMsgRegisterPairs(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgPlaceOrder          = "op_weight_msg_dex"
		defaultWeightMsgPlaceOrder int = 100
	)

	var weightMsgPlaceOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgPlaceOrder, &weightMsgPlaceOrder, nil,
		func(_ *rand.Rand) {
			weightMsgPlaceOrder = defaultWeightMsgPlaceOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlaceOrder,
		dexsimulation.SimulateMsgPlaceOrder(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCleanOrders          = "op_weight_msg_dex"
		defaultWeightMsgCleanOrders int = 100
	)

	var weightMsgCleanOrders int
	simState.AppParams.GetOrGenerate(opWeightMsgCleanOrders, &weightMsgCleanOrders, nil,
		func(_ *rand.Rand) {
			weightMsgCleanOrders = defaultWeightMsgCleanOrders
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCleanOrders,
		dexsimulation.SimulateMsgCleanOrders(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
