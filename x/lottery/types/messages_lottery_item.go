package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateLotteryItem = "create_lottery_item"
	TypeMsgUpdateLotteryItem = "update_lottery_item"
	TypeMsgDeleteLotteryItem = "delete_lottery_item"
)

var _ sdk.Msg = &MsgCreateLotteryItem{}

func NewMsgCreateLotteryItem(creator string, amount int32, participants []string) *MsgCreateLotteryItem {
	return &MsgCreateLotteryItem{
		Creator:      creator,
		Amount:       amount,
		Participants: participants,
	}
}

func (msg *MsgCreateLotteryItem) Route() string {
	return RouterKey
}

func (msg *MsgCreateLotteryItem) Type() string {
	return TypeMsgCreateLotteryItem
}

func (msg *MsgCreateLotteryItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLotteryItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLotteryItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateLotteryItem{}

func NewMsgUpdateLotteryItem(creator string, amount int32, participants []string) *MsgUpdateLotteryItem {
	return &MsgUpdateLotteryItem{
		Creator:      creator,
		Amount:       amount,
		Participants: participants,
	}
}

func (msg *MsgUpdateLotteryItem) Route() string {
	return RouterKey
}

func (msg *MsgUpdateLotteryItem) Type() string {
	return TypeMsgUpdateLotteryItem
}

func (msg *MsgUpdateLotteryItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateLotteryItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateLotteryItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteLotteryItem{}

func NewMsgDeleteLotteryItem(creator string) *MsgDeleteLotteryItem {
	return &MsgDeleteLotteryItem{
		Creator: creator,
	}
}
func (msg *MsgDeleteLotteryItem) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLotteryItem) Type() string {
	return TypeMsgDeleteLotteryItem
}

func (msg *MsgDeleteLotteryItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteLotteryItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteLotteryItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
