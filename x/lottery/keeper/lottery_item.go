package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery/x/lottery/types"
)

// SetLotteryItem set lotteryItem in the store
func (k Keeper) SetLotteryItem(ctx sdk.Context, lotteryItem types.LotteryItem) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryItemKey))
	b := k.cdc.MustMarshal(&lotteryItem)
	store.Set([]byte{0}, b)
}

// GetLotteryItem returns lotteryItem
func (k Keeper) GetLotteryItem(ctx sdk.Context) (val types.LotteryItem, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryItemKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLotteryItem removes lotteryItem from the store
func (k Keeper) RemoveLotteryItem(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryItemKey))
	store.Delete([]byte{0})
}
