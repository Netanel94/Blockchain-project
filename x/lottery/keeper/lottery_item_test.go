package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "lottery/testutil/keeper"
	"lottery/testutil/nullify"
	"lottery/x/lottery/keeper"
	"lottery/x/lottery/types"
)

func createTestLotteryItem(keeper *keeper.Keeper, ctx sdk.Context) types.LotteryItem {
	item := types.LotteryItem{}
	keeper.SetLotteryItem(ctx, item)
	return item
}

func TestLotteryItemGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestLotteryItem(keeper, ctx)
	rst, found := keeper.GetLotteryItem(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLotteryItemRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestLotteryItem(keeper, ctx)
	keeper.RemoveLotteryItem(ctx)
	_, found := keeper.GetLotteryItem(ctx)
	require.False(t, found)
}
