package keeper_test

import (
	"testing"

	"ob/x/dex/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:    types.DefaultParams(),
		MarketMap: []types.Market{{Index: "0"}, {Index: "1"}}, OrderMap: []types.Order{{Index: "0"}, {Index: "1"}}, OrderbookMap: []types.Orderbook{{Index: "0"}, {Index: "1"}}}

	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.MarketMap, got.MarketMap)
	require.EqualExportedValues(t, genesisState.OrderMap, got.OrderMap)
	require.EqualExportedValues(t, genesisState.OrderbookMap, got.OrderbookMap)

}
