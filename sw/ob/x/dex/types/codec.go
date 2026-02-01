package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCleanOrders{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPlaceOrder{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterPairs{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrder{},
		&MsgUpdateOrder{},
		&MsgDeleteOrder{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMarket{},
		&MsgUpdateMarket{},
		&MsgDeleteMarket{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
