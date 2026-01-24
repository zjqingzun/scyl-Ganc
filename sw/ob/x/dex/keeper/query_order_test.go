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

func createNOrder(keeper keeper.Keeper, ctx context.Context, n int) []types.Order {
	items := make([]types.Order, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].MarketId = strconv.Itoa(i)
		items[i].OrderType = strconv.Itoa(i)
		items[i].Side = strconv.Itoa(i)
		items[i].Price = strconv.Itoa(i)
		items[i].Quantity = strconv.Itoa(i)
		items[i].Remaining = strconv.Itoa(i)
		items[i].CreatedAt = strconv.Itoa(i)
		items[i].CreatedHeight = strconv.Itoa(i)
		items[i].Status = strconv.Itoa(i)
		_ = keeper.Order.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestOrderQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNOrder(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetOrderRequest
		response *types.QueryGetOrderResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetOrderRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetOrderResponse{Order: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetOrderRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetOrderResponse{Order: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetOrderRequest{
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
			response, err := qs.GetOrder(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestOrderQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNOrder(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOrderRequest {
		return &types.QueryAllOrderRequest{
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
			resp, err := qs.ListOrder(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Order), step)
			require.Subset(t, msgs, resp.Order)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListOrder(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Order), step)
			require.Subset(t, msgs, resp.Order)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListOrder(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Order)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListOrder(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
