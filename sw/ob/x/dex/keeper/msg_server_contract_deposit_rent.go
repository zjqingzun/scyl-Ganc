package keeper

import (
	"context"

	"ob/x/dex/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) ContractDepositRent(ctx context.Context, msg *types.MsgContractDepositRent) (*types.MsgContractDepositRentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgContractDepositRentResponse{}, nil
}
