package keeper

import (
	"context"
	"encoding/hex"
	"strconv"

	"rps/x/rps/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/crypto/sha3"
)

func (k msgServer) RevealGame(goCtx context.Context, msg *types.MsgRevealGame) (*types.MsgRevealGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	game, found := k.GetGames(ctx, strconv.FormatUint(msg.GameId, 10))
	if !found {
		return nil, types.NoGameFoundError
	}

	if (game.Player1 != msg.Creator) && (game.Player2 != msg.Creator) {
		return nil, types.NotPlayerError
	}
	if game.Player2 == "" {
		return nil, types.GameStatusError
	}
	turnHash := ""
	if game.Player1 == msg.Creator {
		turnHash = game.TurnHash1
	} else {
		turnHash = game.TurnHash2
	}
	msgTurnHashBytes := sha3.Sum224([]byte(msg.Turn + msg.Salt))
	msgTurnHash := hex.EncodeToString(msgTurnHashBytes[:])
	if turnHash != msgTurnHash {
		return nil, types.WrongSaltError
	}
	if game.Player1 == msg.Creator {
		game.Turn1 = msg.Turn
	} else {
		game.Turn2 = msg.Turn
	}
	k.SetGames(ctx, game)
	if (game.Turn1 != "") && (game.Turn2 != "") {
		result := ResolveGame(game)
		totalWin := sdk.Coin{Denom: game.BetAmount.Denom, Amount: game.BetAmount.Amount.MulRaw(2)}
		player1Addr, err := sdk.AccAddressFromBech32(game.Player1)
		if err != nil {
			panic(err)
		}
		player2Addr, err := sdk.AccAddressFromBech32(game.Player2)
		if err != nil {
			panic(err)
		}
		if result == Player1 {
			sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, player1Addr, sdk.Coins{totalWin})
			if sdkError != nil {
				return nil, sdkError
			}
		} else if result == Player2 {
			sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, player2Addr, sdk.Coins{totalWin})
			if sdkError != nil {
				return nil, sdkError
			}
		} else {
			sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, player1Addr, sdk.Coins{game.BetAmount})
			if sdkError != nil {
				return nil, sdkError
			}
			sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, player2Addr, sdk.Coins{game.BetAmount})
			if sdkError != nil {
				return nil, sdkError
			}
		}
		k.RemoveGames(ctx, strconv.FormatUint(msg.GameId, 10))
	}
	return &types.MsgRevealGameResponse{}, nil
}

type GameResult uint8

const (
	Draw GameResult = iota
	Player1
	Player2
)

func ResolveGame(game types.Games) GameResult {
	if game.Turn1 == game.Turn2 {
		return Draw
	} else if game.Turn1 == "rock" {
		if game.Turn2 == "scissors" {
			return Player1
		} else {
			return Player2
		}
	} else if game.Turn1 == "paper" {
		if game.Turn2 == "rock" {
			return Player1
		} else {
			return Player2
		}
	} else if game.Turn1 == "scissors" {
		if game.Turn2 == "paper" {
			return Player1
		} else {
			return Player2
		}
	}
	return Draw
}
