package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"ob/x/dex/keeper"
	"ob/x/dex/types"
)

func createNOrderbook(keeper keeper.Keeper, ctx context.Context, n int) []types.Orderbook {
	items := make([]types.Orderbook, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].MarketId = strconv.Itoa(i)
		items[i].Side = strconv.Itoa(i)
		items[i].Price = strconv.Itoa(i)
		items[i].OrderId = strconv.Itoa(i)
		_ = keeper.Orderbook.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestOrderbookQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNOrderbook(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetOrderbookRequest
		response *types.QueryGetOrderbookResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetOrderbookRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetOrderbookResponse{Orderbook: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetOrderbookRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetOrderbookResponse{Orderbook: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetOrderbookRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetOrderbook(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestOrderbookQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNOrderbook(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOrderbookRequest {
		return &types.QueryAllOrderbookRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListOrderbook(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Orderbook), step)
			require.Subset(t, msgs, resp.Orderbook)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListOrderbook(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Orderbook), step)
			require.Subset(t, msgs, resp.Orderbook)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListOrderbook(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Orderbook)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListOrderbook(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
