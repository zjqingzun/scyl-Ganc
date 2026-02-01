package keeper

import (
	"context"

	"ob/x/dex/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CleanOrders(goCtx context.Context, msg *types.MsgCleanOrders) (*types.MsgCleanOrdersResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message
	ctx := sdk.UnwrapSDKContext(goCtx)

	limit := msg.Limit
	if limit > 500 { limit = 500 }
	
	count := k.Keeper.BatchCleanFilledOrders(ctx, limit)

	
	return &types.MsgCleanOrdersResponse{Count: count}, nil
	// return &types.MsgCleanOrdersResponse{}, nil
}
