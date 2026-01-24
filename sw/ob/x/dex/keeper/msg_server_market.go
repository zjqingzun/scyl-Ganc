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

func (k msgServer) CreateMarket(ctx context.Context, msg *types.MsgCreateMarket) (*types.MsgCreateMarketResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Market.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var market = types.Market{
		Creator:    msg.Creator,
		Index:      msg.Index,
		BaseDenom:  msg.BaseDenom,
		QuoteDenom: msg.QuoteDenom,
		TickSize:   msg.TickSize,
		LotSize:    msg.LotSize,
		Status:     msg.Status,
	}

	if err := k.Market.Set(ctx, market.Index, market); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateMarketResponse{}, nil
}

func (k msgServer) UpdateMarket(ctx context.Context, msg *types.MsgUpdateMarket) (*types.MsgUpdateMarketResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Market.Get(ctx, msg.Index)
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

	var market = types.Market{
		Creator:    msg.Creator,
		Index:      msg.Index,
		BaseDenom:  msg.BaseDenom,
		QuoteDenom: msg.QuoteDenom,
		TickSize:   msg.TickSize,
		LotSize:    msg.LotSize,
		Status:     msg.Status,
	}

	if err := k.Market.Set(ctx, market.Index, market); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update market")
	}

	return &types.MsgUpdateMarketResponse{}, nil
}

func (k msgServer) DeleteMarket(ctx context.Context, msg *types.MsgDeleteMarket) (*types.MsgDeleteMarketResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Market.Get(ctx, msg.Index)
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

	if err := k.Market.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove market")
	}

	return &types.MsgDeleteMarketResponse{}, nil
}
