package keeper

import (
	"context"
	"errors"
	"fmt"

	"ob/x/dex/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateOrder(ctx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Order.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var order = types.Order{
		Creator:       msg.Creator,
		Index:         msg.Index,
		MarketId:      msg.MarketId,
		OrderType:     msg.OrderType,
		Side:          msg.Side,
		Price:         msg.Price,
		Quantity:      msg.Quantity,
		Remaining:     msg.Remaining,
		CreatedAt:     msg.CreatedAt,
		CreatedHeight: msg.CreatedHeight,
		Status:        msg.Status,
	}

	if err := k.Order.Set(ctx, order.Index, order); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateOrderResponse{}, nil
}

func (k msgServer) UpdateOrder(ctx context.Context, msg *types.MsgUpdateOrder) (*types.MsgUpdateOrderResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Order.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var order = types.Order{
		Creator:       msg.Creator,
		Index:         msg.Index,
		MarketId:      msg.MarketId,
		OrderType:     msg.OrderType,
		Side:          msg.Side,
		Price:         msg.Price,
		Quantity:      msg.Quantity,
		Remaining:     msg.Remaining,
		CreatedAt:     msg.CreatedAt,
		CreatedHeight: msg.CreatedHeight,
		Status:        msg.Status,
	}

	if err := k.Order.Set(ctx, order.Index, order); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update order")
	}

	return &types.MsgUpdateOrderResponse{}, nil
}

func (k msgServer) DeleteOrder(ctx context.Context, msg *types.MsgDeleteOrder) (*types.MsgDeleteOrderResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Order.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Order.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove order")
	}

	return &types.MsgDeleteOrderResponse{}, nil
}
