package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"rps/x/rps/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateGameResponse{}, nil
}
