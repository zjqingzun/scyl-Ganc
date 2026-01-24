package keeper

import (
	"context"
	"errors"

	"ob/x/dex/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListOrderbook(ctx context.Context, req *types.QueryAllOrderbookRequest) (*types.QueryAllOrderbookResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    // Nếu bạn muốn lọc chỉ lấy của 1 Market (giả sử req có field MarketId)
    // Bạn có thể dùng q.k.Orderbook.Iterate với Prefix.
    // Nhưng hiện tại để test list-all, logic của bạn đã ổn.

    orderbooks, pageRes, err := query.CollectionPaginate(
        ctx,
        q.k.Orderbook,
        req.Pagination,
        func(key string, value types.Orderbook) (types.Orderbook, error) {
            return value, nil
        },
    )
    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryAllOrderbookResponse{Orderbook: orderbooks, Pagination: pageRes}, nil
}

func (q queryServer) GetOrderbook(ctx context.Context, req *types.QueryGetOrderbookRequest) (*types.QueryGetOrderbookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Orderbook.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetOrderbookResponse{Orderbook: val}, nil
}
