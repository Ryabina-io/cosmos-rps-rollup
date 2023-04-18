package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rps/x/rps/types"
)

func (k Keeper) GamesAll(goCtx context.Context, req *types.QueryAllGamesRequest) (*types.QueryAllGamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var gamess []types.Games
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	gamesStore := prefix.NewStore(store, types.KeyPrefix(types.GamesKeyPrefix))

	pageRes, err := query.Paginate(gamesStore, req.Pagination, func(key []byte, value []byte) error {
		var games types.Games
		if err := k.cdc.Unmarshal(value, &games); err != nil {
			return err
		}

		gamess = append(gamess, games)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGamesResponse{Games: gamess, Pagination: pageRes}, nil
}

func (k Keeper) Games(goCtx context.Context, req *types.QueryGetGamesRequest) (*types.QueryGetGamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetGames(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGamesResponse{Games: val}, nil
}
