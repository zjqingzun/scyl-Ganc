package keeper

import (
	"context"
	"ob/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterPairs(goCtx context.Context, msg *types.MsgRegisterPairs) (*types.MsgRegisterPairsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	marketId := msg.BaseDenom + "-" + msg.QuoteDenom

	market := types.Market{
		BaseDenom:  msg.BaseDenom,
		QuoteDenom: msg.QuoteDenom,
		TickSize:   msg.TickSize,
		LotSize:    msg.LotSize,
		Status:     "ACTIVE",
	}

	// Sử dụng k.Keeper
	if err := k.Keeper.Market.Set(ctx, marketId, market); err != nil {
		return nil, err
	}

	return &types.MsgRegisterPairsResponse{Id: marketId}, nil
}