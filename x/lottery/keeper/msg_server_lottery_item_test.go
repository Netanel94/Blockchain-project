package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "lottery/testutil/keeper"
	"lottery/x/lottery/keeper"
	"lottery/x/lottery/types"
)

func TestLotteryItemMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateLotteryItem{Creator: creator}
	_, err := srv.CreateLotteryItem(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetLotteryItem(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestLotteryItemMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateLotteryItem
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateLotteryItem{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateLotteryItem{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LotteryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateLotteryItem{Creator: creator}
			_, err := srv.CreateLotteryItem(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateLotteryItem(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetLotteryItem(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestLotteryItemMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteLotteryItem
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteLotteryItem{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteLotteryItem{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LotteryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateLotteryItem(wctx, &types.MsgCreateLotteryItem{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteLotteryItem(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetLotteryItem(ctx)
				require.False(t, found)
			}
		})
	}
}
