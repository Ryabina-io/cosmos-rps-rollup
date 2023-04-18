package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "rps/testutil/keeper"
	"rps/x/rps/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RpsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
