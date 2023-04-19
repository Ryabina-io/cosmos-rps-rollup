package keeper

import (
	"context"
	"strconv"

	"rps/x/rps/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	info, _ := k.GetSystemInfo(ctx)
	game := types.Games{
		Index:     strconv.FormatUint(info.NextId, 10),
		BetAmount: msg.BetAmount,
		Player1:   msg.Creator,
		Player2:   "",
		TurnHash1: msg.TurnHash,
		TurnHash2: "",
		Turn1:     "",
		Turn2:     "",
	}
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{msg.BetAmount})
	if sdkError != nil {
		return nil, sdkError
	}
	k.SetGames(ctx, game)
	info.NextId++
	k.SetSystemInfo(ctx, info)

	return &types.MsgCreateGameResponse{
		GameId: game.Index,
	}, nil
}
