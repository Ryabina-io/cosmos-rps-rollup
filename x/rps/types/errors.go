package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rps module sentinel errors
var (
	NoGameFoundError       = sdkerrors.Register(ModuleName, 1100, "No Game Found")
	GameAlreadyJoinedError = sdkerrors.Register(ModuleName, 1101, "Game Already Joined")
	NotGameOwnerError      = sdkerrors.Register(ModuleName, 1102, "Not Game Owner")
)
