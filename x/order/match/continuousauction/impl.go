package continuousauction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/omexapp/omexchain/x/order/keeper"
)

// nolint
type CaEngine struct {
}

// nolint
func (e *CaEngine) Run(ctx sdk.Context, keeper keeper.Keeper) {
	// TODO
}
