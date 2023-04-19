package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveGame = "remove_game"

var _ sdk.Msg = &MsgRemoveGame{}

func NewMsgRemoveGame(creator string, gameId uint64) *MsgRemoveGame {
	return &MsgRemoveGame{
		Creator: creator,
		GameId:  gameId,
	}
}

func (msg *MsgRemoveGame) Route() string {
	return RouterKey
}

func (msg *MsgRemoveGame) Type() string {
	return TypeMsgRemoveGame
}

func (msg *MsgRemoveGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
