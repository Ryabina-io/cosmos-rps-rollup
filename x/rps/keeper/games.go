package keeper

import (
	"rps/x/rps/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetGames set a specific games in the store from its index
func (k Keeper) SetGames(ctx sdk.Context, games types.Games) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GamesKeyPrefix))
	b := k.cdc.MustMarshal(&games)
	store.Set(types.GamesKey(
		games.Index,
	), b)
}

// GetGames returns a games from its index
func (k Keeper) GetGames(
	ctx sdk.Context,
	index string,

) (val types.Games, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GamesKeyPrefix))

	b := store.Get(types.GamesKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGames removes a games from the store
func (k Keeper) RemoveGames(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GamesKeyPrefix))
	store.Delete(types.GamesKey(
		index,
	))
}

// GetAllGames returns all games
func (k Keeper) GetAllGames(ctx sdk.Context) (list []types.Games) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GamesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Games
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
