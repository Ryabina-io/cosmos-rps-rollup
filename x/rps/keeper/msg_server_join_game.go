package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"rps/x/rps/types"
)

func (k msgServer) JoinGame(goCtx context.Context, msg *types.MsgJoinGame) (*types.MsgJoinGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgJoinGameResponse{}, nil
}
