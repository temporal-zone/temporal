package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/temporal-zone/temporal/app"
)

func initSDKConfig() {
	// Set prefixes
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	// Set the DefaultBondDenom
	sdk.DefaultBondDenom = app.DefaultBondDenom

	// Register token denom
	err := sdk.RegisterDenom(app.HumanCoinUnit, sdk.OneDec())
	if err != nil {
		panic(err)
	}

	err = sdk.RegisterDenom(app.BaseCoinUnit, sdk.NewDecWithPrec(1, app.TprlExponent))
	if err != nil {
		panic(err)
	}

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}
