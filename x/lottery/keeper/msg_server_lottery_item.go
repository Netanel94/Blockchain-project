package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"lottery/x/lottery/types"
)

func (k msgServer) CreateLotteryItem(goCtx context.Context, msg *types.MsgCreateLotteryItem) (*types.MsgCreateLotteryItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetLotteryItem(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var lotteryItem = types.LotteryItem{
		Creator:      msg.Creator,
		Amount:       msg.Amount,
		Participants: msg.Participants,
	}

	k.SetLotteryItem(
		ctx,
		lotteryItem,
	)
	return &types.MsgCreateLotteryItemResponse{}, nil
}

func (k msgServer) UpdateLotteryItem(goCtx context.Context, msg *types.MsgUpdateLotteryItem) (*types.MsgUpdateLotteryItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLotteryItem(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var lotteryItem = types.LotteryItem{
		Creator:      msg.Creator,
		Amount:       msg.Amount,
		Participants: msg.Participants,
	}

	k.SetLotteryItem(ctx, lotteryItem)

	return &types.MsgUpdateLotteryItemResponse{}, nil
}

func (k msgServer) DeleteLotteryItem(goCtx context.Context, msg *types.MsgDeleteLotteryItem) (*types.MsgDeleteLotteryItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLotteryItem(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLotteryItem(ctx)

	return &types.MsgDeleteLotteryItemResponse{}, nil
}
