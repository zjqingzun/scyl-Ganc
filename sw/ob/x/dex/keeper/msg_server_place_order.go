package keeper

import (
	"context"
	"fmt"

	"ob/x/dex/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PlaceOrder(goCtx context.Context, msg *types.MsgPlaceOrder) (*types.MsgPlaceOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := k.Keeper.Market.Get(ctx, msg.MarketId)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "market %s not found", msg.MarketId)
	}

	orderType := "LIMIT"
	if msg.Price == "0" {
		orderType = "MARKET"
	}

	orderId := fmt.Sprintf("%s-%d", msg.MarketId, ctx.BlockHeight())
	order := types.Order{
		MarketId:      msg.MarketId,
		OrderType:     orderType,
		Side:          msg.Side,
		Price:         msg.Price,
		Quantity:      msg.Quantity,
		Remaining:     msg.Quantity,
		CreatedAt:     fmt.Sprintf("%d", ctx.BlockTime().Unix()),
		CreatedHeight: fmt.Sprintf("%d", ctx.BlockHeight()),
		Status:        "OPEN",
		Creator:       msg.Creator,
	}

	// Save via Keeper's Collections API
	if err := k.Keeper.Order.Set(ctx, orderId, order); err != nil {
		return nil, err
	}

	// Create a composite key to automatically sort by: marketId | price | orderId
	if orderType == "LIMIT" {
		orderbookKey := fmt.Sprintf("%s|%s|%s", msg.MarketId, msg.Price, orderId)
		err = k.Keeper.Orderbook.Set(ctx, orderbookKey, types.Orderbook{
			MarketId: msg.MarketId,
			Side:     msg.Side,
			Price:    msg.Price,
			OrderId:  orderId,
		})
		if err != nil {
			return nil, err
		}
	}

	// ********** Matching Order **********
	err = k.Keeper.MatchOrders(ctx, &order, orderId)
	if err != nil {
		return nil, err
	}

	return &types.MsgPlaceOrderResponse{OrderId: orderId}, nil
}