package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateUserInstruction{}, "record/CreateUserInstruction", nil)
	cdc.RegisterConcrete(&MsgDeleteUserInstruction{}, "record/DeleteUserInstruction", nil)
	cdc.RegisterConcrete(&MsgUpdateUserInstruction{}, "record/UpdateUserInstruction", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUserInstruction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteUserInstruction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateUserInstruction{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
