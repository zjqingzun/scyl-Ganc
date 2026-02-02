package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BatchCleanFilledOrders scans and deletes FILLED orders with limit restrictions.
func (k Keeper) BatchCleanFilledOrders(ctx sdk.Context, limit uint64) uint64 {
	var deletedCount uint64 = 0

	iter, err := k.Order.Iterate(ctx, nil)
	if err != nil {
		return 0
	}
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		if deletedCount >= limit {
			break
		}

		order, err := iter.Value()
		if err != nil {
			continue
		}

		if order.Status == "FILLED" {
			orderId, err := iter.Key()
			if err != nil {
				continue
			}

			k.Order.Remove(ctx, orderId)

			orderbookKey := fmt.Sprintf("%s|%s|%s", order.MarketId, order.Price, orderId)
			k.Orderbook.Remove(ctx, orderbookKey)

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					"order_purged",
					sdk.NewAttribute("order_id", orderId),
					sdk.NewAttribute("creator", order.Creator),
				),
			)

			deletedCount++
		}
	}

	if deletedCount > 0 {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"batch_cleanup",
				sdk.NewAttribute("module", "dex"),
				sdk.NewAttribute("deleted_count", fmt.Sprintf("%d", deletedCount)),
			),
		)

		ctx.Logger().Info("DEX Cleanup", "deleted", deletedCount)
	}

	return deletedCount
}
