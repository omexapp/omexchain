package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/omexapp/omexchain/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestModuleAccountInvariant(t *testing.T) {

	testInput := createTestInputWithBalance(t, 1, 10000)
	ctx := testInput.Ctx
	keeper := testInput.DexKeeper
	accounts := testInput.TestAddrs
	keeper.SetParams(ctx, *types.DefaultParams())

	builtInTP := GetBuiltInTokenPair()
	builtInTP.Owner = accounts[0]
	err := keeper.SaveTokenPair(ctx, builtInTP)
	require.Nil(t, err)

	// deposit xxb_omt 100 omt
	depositMsg := types.NewMsgDeposit(builtInTP.Name(),
		sdk.NewDecCoin(builtInTP.QuoteAssetSymbol, sdk.NewInt(100)), accounts[0])

	err = keeper.Deposit(ctx, builtInTP.Name(), depositMsg.Depositor, depositMsg.Amount)
	require.Nil(t, err)

	// module acount balance 100omt
	// xxb_omt deposits 100 omt. withdraw info 0 omt
	invariant := ModuleAccountInvariant(keeper, keeper.supplyKeeper)
	_, broken := invariant(ctx)
	require.False(t, broken)

	// withdraw xxb_omt 50 omt
	WithdrawMsg := types.NewMsgWithdraw(builtInTP.Name(),
		sdk.NewDecCoin(builtInTP.QuoteAssetSymbol, sdk.NewInt(50)), accounts[0])

	err = keeper.Withdraw(ctx, builtInTP.Name(), WithdrawMsg.Depositor, WithdrawMsg.Amount)
	require.Nil(t, err)

	// module acount balance 100omt
	// xxb_omt deposits 50 omt. withdraw info 50 omt
	invariant = ModuleAccountInvariant(keeper, keeper.supplyKeeper)
	_, broken = invariant(ctx)
	require.False(t, broken)

}
