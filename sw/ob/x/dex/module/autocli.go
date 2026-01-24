package dex

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"ob/x/dex/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListMarket",
					Use:       "list-market",
					Short:     "List all market",
				},
				{
					RpcMethod:      "GetMarket",
					Use:            "get-market [id]",
					Short:          "Gets a market",
					Alias:          []string{"show-market"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListOrder",
					Use:       "list-order",
					Short:     "List all order",
				},
				{
					RpcMethod:      "GetOrder",
					Use:            "get-order [id]",
					Short:          "Gets a order",
					Alias:          []string{"show-order"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListOrderbook",
					Use:       "list-orderbook",
					Short:     "List all orderbook",
				},
				{
					RpcMethod:      "GetOrderbook",
					Use:            "get-orderbook [id]",
					Short:          "Gets a orderbook",
					Alias:          []string{"show-orderbook"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateMarket",
					Use:            "create-market [index] [base-denom] [quote-denom] [tick-size] [lot-size] [status]",
					Short:          "Create a new market",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "base_denom"}, {ProtoField: "quote_denom"}, {ProtoField: "tick_size"}, {ProtoField: "lot_size"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "UpdateMarket",
					Use:            "update-market [index] [base-denom] [quote-denom] [tick-size] [lot-size] [status]",
					Short:          "Update market",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "base_denom"}, {ProtoField: "quote_denom"}, {ProtoField: "tick_size"}, {ProtoField: "lot_size"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeleteMarket",
					Use:            "delete-market [index]",
					Short:          "Delete market",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateOrder",
					Use:            "create-order [index] [market-id] [order-type] [side] [price] [quantity] [remaining] [created-at] [created-height] [status]",
					Short:          "Create a new order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "market_id"}, {ProtoField: "order_type"}, {ProtoField: "side"}, {ProtoField: "price"}, {ProtoField: "quantity"}, {ProtoField: "remaining"}, {ProtoField: "created_at"}, {ProtoField: "created_height"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "UpdateOrder",
					Use:            "update-order [index] [market-id] [order-type] [side] [price] [quantity] [remaining] [created-at] [created-height] [status]",
					Short:          "Update order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "market_id"}, {ProtoField: "order_type"}, {ProtoField: "side"}, {ProtoField: "price"}, {ProtoField: "quantity"}, {ProtoField: "remaining"}, {ProtoField: "created_at"}, {ProtoField: "created_height"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeleteOrder",
					Use:            "delete-order [index]",
					Short:          "Delete order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "RegisterPairs",
					Use:            "register-pairs [base-denom] [quote-denom] [tick-size] [lot-size]",
					Short:          "Send a registerPairs tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "base_denom"}, {ProtoField: "quote_denom"}, {ProtoField: "tick_size"}, {ProtoField: "lot_size"}},
				},
				{
					RpcMethod:      "PlaceOrder",
					Use:            "place-order [market-id] [side] [price] [quantity]",
					Short:          "Send a placeOrder tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "market_id"}, {ProtoField: "side"}, {ProtoField: "price"}, {ProtoField: "quantity"}},
				},
				{
					RpcMethod:      "CancelOrder",
					Use:            "cancel-order [order-id]",
					Short:          "Send a cancelOrder tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "order_id"}},
				},
				{
					RpcMethod:      "RegisterContract",
					Use:            "register-contract [market-id] [contract-addr]",
					Short:          "Send a registerContract tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "market_id"}, {ProtoField: "contract_addr"}},
				},
				{
					RpcMethod:      "UnregisterContract",
					Use:            "unregister-contract [contract-addr]",
					Short:          "Send a unregisterContract tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "contract_addr"}},
				},
				{
					RpcMethod:      "ContractDepositRent",
					Use:            "contract-deposit-rent [contract-addr] [amount]",
					Short:          "Send a contractDepositRent tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "contract_addr"}, {ProtoField: "amount"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
