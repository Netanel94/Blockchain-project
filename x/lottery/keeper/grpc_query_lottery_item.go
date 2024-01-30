package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lottery/x/lottery/types"
)

func (k Keeper) LotteryItem(c context.Context, req *types.QueryGetLotteryItemRequest) (*types.QueryGetLotteryItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLotteryItem(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLotteryItemResponse{LotteryItem: val}, nil
}
