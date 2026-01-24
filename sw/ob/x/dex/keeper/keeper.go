package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"ob/x/dex/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema collections.Schema
	Params collections.Item[types.Params]

	authKeeper types.AuthKeeper
	bankKeeper types.BankKeeper
	Market     collections.Map[string, types.Market]
	Order      collections.Map[string, types.Order]
	Orderbook  collections.Map[string, types.Orderbook]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,

	authKeeper types.AuthKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		authKeeper: authKeeper,
		bankKeeper: bankKeeper,
		Params:     collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Market:     collections.NewMap(sb, types.MarketKey, "market", collections.StringKey, codec.CollValue[types.Market](cdc)), Order: collections.NewMap(sb, types.OrderKey, "order", collections.StringKey, codec.CollValue[types.Order](cdc)), Orderbook: collections.NewMap(sb, types.OrderbookKey, "orderbook", collections.StringKey, codec.CollValue[types.Orderbook](cdc))}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}
