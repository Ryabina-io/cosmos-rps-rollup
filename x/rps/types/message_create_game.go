package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGame = "create_game"

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(creator string, betAmount sdk.Coin, turnHash string) *MsgCreateGame {
	return &MsgCreateGame{
		Creator:   creator,
		BetAmount: betAmount,
		TurnHash:  turnHash,
	}
}

func (msg *MsgCreateGame) Route() string {
	return RouterKey
}

func (msg *MsgCreateGame) Type() string {
	return TypeMsgCreateGame
}

func (msg *MsgCreateGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
