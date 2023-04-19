package keeper

import (
	"context"
	"strconv"

	"rps/x/rps/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) JoinGame(goCtx context.Context, msg *types.MsgJoinGame) (*types.MsgJoinGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	game, found := k.GetGames(ctx, strconv.FormatUint(msg.GameId, 10))
	if !found {
		return nil, types.NoGameFoundError
	}
	if game.Player2 != "" {
		return nil, types.GameAlreadyJoinedError
	}
	if game.Player1 == msg.Creator {
		return nil, types.GameAlreadyJoinedError
	}
	game.Player2 = msg.Creator
	game.TurnHash2 = msg.TurnHash
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{game.BetAmount})
	if sdkError != nil {
		return nil, sdkError
	}
	k.SetGames(ctx, game)
	return &types.MsgJoinGameResponse{}, nil
}
