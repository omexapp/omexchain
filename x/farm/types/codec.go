package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePool{}, "omexchain/farm/MsgCreatePool", nil)
	cdc.RegisterConcrete(MsgDestroyPool{}, "omexchain/farm/MsgDestroyPool", nil)
	cdc.RegisterConcrete(MsgLock{}, "omexchain/farm/MsgLock", nil)
	cdc.RegisterConcrete(MsgUnlock{}, "omexchain/farm/MsgUnlock", nil)
	cdc.RegisterConcrete(MsgClaim{}, "omexchain/farm/MsgClaim", nil)
	cdc.RegisterConcrete(MsgProvide{}, "omexchain/farm/MsgProvide", nil)
	cdc.RegisterConcrete(ManageWhiteListProposal{}, "omexchain/farm/ManageWhiteListProposal", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
