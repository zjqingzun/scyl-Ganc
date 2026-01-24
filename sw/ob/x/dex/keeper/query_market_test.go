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

func createNMarket(keeper keeper.Keeper, ctx context.Context, n int) []types.Market {
	items := make([]types.Market, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].BaseDenom = strconv.Itoa(i)
		items[i].QuoteDenom = strconv.Itoa(i)
		items[i].TickSize = strconv.Itoa(i)
		items[i].LotSize = strconv.Itoa(i)
		items[i].Status = strconv.Itoa(i)
		_ = keeper.Market.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestMarketQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNMarket(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetMarketRequest
		response *types.QueryGetMarketResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMarketRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetMarketResponse{Market: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMarketRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetMarketResponse{Market: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMarketRequest{
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
			response, err := qs.GetMarket(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestMarketQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNMarket(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMarketRequest {
		return &types.QueryAllMarketRequest{
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
			resp, err := qs.ListMarket(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Market), step)
			require.Subset(t, msgs, resp.Market)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListMarket(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Market), step)
			require.Subset(t, msgs, resp.Market)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListMarket(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Market)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListMarket(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
