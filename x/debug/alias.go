package debug

import (
	"github.com/omexapp/omexchain/x/debug/keeper"
	"github.com/omexapp/omexchain/x/debug/types"
)

const (
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	ModuleName   = types.ModuleName
)

var (
	// functions aliases
	RegisterCodec  = types.RegisterCodec
	NewDebugKeeper = keeper.NewDebugKeeper
	NewDebugger    = keeper.NewDebugger
)

type (
	Keeper = keeper.Keeper
)
