package rps_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "rps/testutil/keeper"
	"rps/testutil/nullify"
	"rps/x/rps"
	"rps/x/rps/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RpsKeeper(t)
	rps.InitGenesis(ctx, *k, genesisState)
	got := rps.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
