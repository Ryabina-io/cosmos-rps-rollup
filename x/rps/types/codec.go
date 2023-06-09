package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGame{}, "rps/CreateGame", nil)
	cdc.RegisterConcrete(&MsgJoinGame{}, "rps/JoinGame", nil)
	cdc.RegisterConcrete(&MsgRevealGame{}, "rps/RevealGame", nil)
	cdc.RegisterConcrete(&MsgRemoveGame{}, "rps/RemoveGame", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGame{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgJoinGame{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevealGame{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveGame{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
