package keeper

import (
	"context"

	"ob/x/dex/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) UnregisterContract(ctx context.Context, msg *types.MsgUnregisterContract) (*types.MsgUnregisterContractResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgUnregisterContractResponse{}, nil
}
