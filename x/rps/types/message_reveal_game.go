package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevealGame = "reveal_game"

var _ sdk.Msg = &MsgRevealGame{}

func NewMsgRevealGame(creator string, gameId uint64, turn string, salt string) *MsgRevealGame {
	return &MsgRevealGame{
		Creator: creator,
		GameId:  gameId,
		Turn:    turn,
		Salt:    salt,
	}
}

func (msg *MsgRevealGame) Route() string {
	return RouterKey
}

func (msg *MsgRevealGame) Type() string {
	return TypeMsgRevealGame
}

func (msg *MsgRevealGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevealGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevealGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
