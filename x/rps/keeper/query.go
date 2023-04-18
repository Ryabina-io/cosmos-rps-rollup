package keeper

import (
	"rps/x/rps/types"
)

var _ types.QueryServer = Keeper{}
