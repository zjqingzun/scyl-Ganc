package keeper

import (
	"context"

	"ob/x/dex/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.MarketMap {
		if err := k.Market.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.OrderMap {
		if err := k.Order.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.OrderbookMap {
		if err := k.Orderbook.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	if err := k.Market.Walk(ctx, nil, func(_ string, val types.Market) (stop bool, err error) {
		genesis.MarketMap = append(genesis.MarketMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.Order.Walk(ctx, nil, func(_ string, val types.Order) (stop bool, err error) {
		genesis.OrderMap = append(genesis.OrderMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.Orderbook.Walk(ctx, nil, func(_ string, val types.Orderbook) (stop bool, err error) {
		genesis.OrderbookMap = append(genesis.OrderbookMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
