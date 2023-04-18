package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "rps/testutil/keeper"
	"rps/testutil/nullify"
	"rps/x/rps/keeper"
	"rps/x/rps/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGames(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Games {
	items := make([]types.Games, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetGames(ctx, items[i])
	}
	return items
}

func TestGamesGet(t *testing.T) {
	keeper, ctx := keepertest.RpsKeeper(t)
	items := createNGames(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGames(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGamesRemove(t *testing.T) {
	keeper, ctx := keepertest.RpsKeeper(t)
	items := createNGames(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGames(ctx,
			item.Index,
		)
		_, found := keeper.GetGames(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestGamesGetAll(t *testing.T) {
	keeper, ctx := keepertest.RpsKeeper(t)
	items := createNGames(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGames(ctx)),
	)
}
