package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"rps/x/rps/types"
)

func (k msgServer) RevealGame(goCtx context.Context, msg *types.MsgRevealGame) (*types.MsgRevealGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevealGameResponse{}, nil
}
