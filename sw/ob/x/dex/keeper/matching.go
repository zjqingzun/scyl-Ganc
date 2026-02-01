package keeper

import (
	"context"
	//"fmt"
	"sort"

	"ob/x/dex/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MatchOrders(goCtx context.Context, newOrder *types.Order, orderId string) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	targetSide := "SELL"
	if newOrder.Side == "SELL" {
		targetSide = "BUY"
	}

	iter, err := k.Orderbook.Iterate(ctx, nil)
	if err != nil {
		return err
	}
	defer iter.Close()

	var candidateOrders []types.Orderbook
	for ; iter.Valid(); iter.Next() {
		val, err := iter.Value()
		if err != nil {
			continue
		}

		if val.MarketId == newOrder.MarketId && val.Side == targetSide {
			if newOrder.OrderType == "LIMIT" {
				if newOrder.Side == "BUY" && val.Price > newOrder.Price {
					continue
				}
				if newOrder.Side == "SELL" && val.Price < newOrder.Price {
					continue
				}
			}
			candidateOrders = append(candidateOrders, val)
		}
	}

	sort.Slice(candidateOrders, func(i, j int) bool {
		if targetSide == "BUY" {
			return candidateOrders[i].Price > candidateOrders[j].Price
		}
		return candidateOrders[i].Price < candidateOrders[j].Price
	})

	newRemaining, ok := math.NewIntFromString(newOrder.Remaining)
	if !ok {
		newRemaining = math.ZeroInt()
	}

	for _, bookEntry := range candidateOrders {
		if newRemaining.IsZero() {
			break
		}

		makerOrder, _ := k.Order.Get(ctx, bookEntry.OrderId)
		makerRemaining, ok := math.NewIntFromString(makerOrder.Remaining)
		if !ok {
			continue
		}

		matchQty := math.MinInt(newRemaining, makerRemaining)

		newRemaining = newRemaining.Sub(matchQty)
		makerRemaining = makerRemaining.Sub(matchQty)

		makerOrder.Remaining = makerRemaining.String()
		if makerRemaining.IsZero() {
			makerOrder.Status = "FILLED"
			k.Order.Set(ctx, bookEntry.OrderId, makerOrder)
			k.Orderbook.Remove(ctx, bookEntry.MarketId+"|"+bookEntry.Price+"|"+bookEntry.OrderId)
		} else {
			k.Order.Set(ctx, bookEntry.OrderId, makerOrder)
		}
	}

	newOrder.Remaining = newRemaining.String()
	if newRemaining.IsZero() {
		newOrder.Status = "FILLED"
		k.Orderbook.Remove(ctx, newOrder.MarketId+"|"+newOrder.Price+"|"+orderId)
	} else if newOrder.OrderType == "MARKET" {
		newOrder.Status = "CANCELLED_REMAINDER"
		k.Orderbook.Remove(ctx, newOrder.MarketId+"|"+newOrder.Price+"|"+orderId)
	} else {
		newOrder.Status = "PARTIAL"
	}

	k.Order.Set(ctx, orderId, *newOrder)
	return nil
}
