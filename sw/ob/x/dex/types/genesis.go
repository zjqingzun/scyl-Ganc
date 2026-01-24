package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:    DefaultParams(),
		MarketMap: []Market{}, OrderMap: []Order{}, OrderbookMap: []Orderbook{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	marketIndexMap := make(map[string]struct{})

	for _, elem := range gs.MarketMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := marketIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for market")
		}
		marketIndexMap[index] = struct{}{}
	}
	orderIndexMap := make(map[string]struct{})

	for _, elem := range gs.OrderMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := orderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for order")
		}
		orderIndexMap[index] = struct{}{}
	}
	orderbookIndexMap := make(map[string]struct{})

	for _, elem := range gs.OrderbookMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := orderbookIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for orderbook")
		}
		orderbookIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
