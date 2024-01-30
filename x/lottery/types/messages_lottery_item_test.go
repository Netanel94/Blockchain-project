package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"lottery/testutil/sample"
)

func TestMsgCreateLotteryItem_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateLotteryItem
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateLotteryItem{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateLotteryItem{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateLotteryItem_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateLotteryItem
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateLotteryItem{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateLotteryItem{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteLotteryItem_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteLotteryItem
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteLotteryItem{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteLotteryItem{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
