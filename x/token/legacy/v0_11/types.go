package v0_11

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/omexapp/omexchain/x/token/legacy/v0_10"
)

const ModuleName = "token"

type (
	// all state that must be provided in genesis file
	GenesisState struct {
		Params       v0_10.Params     `json:"params"`
		Tokens       []Token          `json:"tokens"`
		LockedAssets []v0_10.AccCoins `json:"locked_assets"`
		LockedFees   []v0_10.AccCoins `json:"locked_fees"`
	}

	Token struct {
		Description         string         `json:"description" v2:"description"`                     // e.g. "OMEX Token"
		Symbol              string         `json:"symbol" v2:"symbol"`                               // e.g. "omt"
		OriginalSymbol      string         `json:"original_symbol" v2:"original_symbol"`             // e.g. "OMT"
		WholeName           string         `json:"whole_name" v2:"whole_name"`                       // e.g. "OMT"
		OriginalTotalSupply sdk.Dec        `json:"original_total_supply" v2:"original_total_supply"` // e.g. 1000000000.00000000
		Owner               sdk.AccAddress `json:"owner" v2:"owner"`                                 // e.g. omexchain1upyg3vl6vqaxqvzts69zpus2c027p7paw63s99
		Mintable            bool           `json:"mintable" v2:"mintable"`                           // e.g. false
	}
)
