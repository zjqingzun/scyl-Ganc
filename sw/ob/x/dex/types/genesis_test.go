package types_test

import (
	"testing"

	"ob/x/dex/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{MarketMap: []types.Market{{Index: "0"}, {Index: "1"}}, OrderMap: []types.Order{{Index: "0"}, {Index: "1"}}, OrderbookMap: []types.Orderbook{{Index: "0"}, {Index: "1"}}},
			valid:    true,
		}, {
			desc: "duplicated market",
			genState: &types.GenesisState{
				MarketMap: []types.Market{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				OrderMap: []types.Order{{Index: "0"}, {Index: "1"}}, OrderbookMap: []types.Orderbook{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated order",
			genState: &types.GenesisState{
				OrderMap: []types.Order{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				OrderbookMap: []types.Orderbook{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated orderbook",
			genState: &types.GenesisState{
				OrderbookMap: []types.Orderbook{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
