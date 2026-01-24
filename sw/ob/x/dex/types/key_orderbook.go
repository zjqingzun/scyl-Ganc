package types

import "cosmossdk.io/collections"

// OrderbookKey is the prefix to retrieve all Orderbook
var OrderbookKey = collections.NewPrefix("orderbook/value/")
