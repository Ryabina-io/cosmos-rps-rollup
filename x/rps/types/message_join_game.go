package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgJoinGame = "join_game"

var _ sdk.Msg = &MsgJoinGame{}

func NewMsgJoinGame(creator string, gameId uint64, turnHash string) *MsgJoinGame {
	return &MsgJoinGame{
		Creator:  creator,
		GameId:   gameId,
		TurnHash: turnHash,
	}
}

func (msg *MsgJoinGame) Route() string {
	return RouterKey
}

func (msg *MsgJoinGame) Type() string {
	return TypeMsgJoinGame
}

func (msg *MsgJoinGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgJoinGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgJoinGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
