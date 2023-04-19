package keeper

import (
	"context"
	"strconv"

	"rps/x/rps/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveGame(goCtx context.Context, msg *types.MsgRemoveGame) (*types.MsgRemoveGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	gameIdStr := strconv.FormatUint(msg.GameId, 10)
	game, found := k.GetGames(ctx, gameIdStr)
	if !found {
		return nil, types.NoGameFoundError
	}
	if game.Player1 != msg.Creator {
		return nil, types.NotGameOwnerError
	}
	if game.Player2 != "" {
		return nil, types.GameAlreadyJoinedError
	}
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{game.BetAmount})
	if sdkError != nil {
		return nil, sdkError
	}
	k.RemoveGames(ctx, gameIdStr)

	return &types.MsgRemoveGameResponse{}, nil
}
