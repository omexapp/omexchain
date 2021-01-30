package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTokenIssue{}, "omexchain/token/MsgIssue", nil)
	cdc.RegisterConcrete(MsgTokenBurn{}, "omexchain/token/MsgBurn", nil)
	cdc.RegisterConcrete(MsgTokenMint{}, "omexchain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "omexchain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgSend{}, "omexchain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "omexchain/token/MsgTransferOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "omexchain/token/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(MsgTokenModify{}, "omexchain/token/MsgModify", nil)

	// for test
	//cdc.RegisterConcrete(MsgTokenDestroy{}, "omexchain/token/MsgDestroy", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
