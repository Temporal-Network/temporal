package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Temporal-Network/temporal/app"
)

const (
	HumanReadableCoinUnit = "TPRL"
	BaseCoinUnit          = "utprl"
	DefaultBondDenom      = BaseCoinUnit
)

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()

	// Register Denoms
	registerDenoms()
}

// RegisterDenoms registers the base and display denominations to the SDK.
func registerDenoms() {
	if err := sdk.RegisterDenom(HumanReadableCoinUnit, sdk.OneDec()); err != nil {
		panic(err)
	}

	if err := sdk.RegisterDenom(BaseCoinUnit, sdk.NewDecWithPrec(1, 6)); err != nil {
		panic(err)
	}
}
