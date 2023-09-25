package apptesting

import (
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/temporal-zone/temporal/app"
)

var (
	ChainID = "temporal-test"
)

type SuitelessAppTestHelper struct {
	App *app.App
	Ctx sdk.Context
}

// Instantiates an TestHelper without the test suite
// This is for testing scenarios where we simply need the setup function to run,
// and need access to the TestHelper attributes and keepers (e.g. genesis tests)
func SetupSuitelessTestHelper() SuitelessAppTestHelper {
	s := SuitelessAppTestHelper{}
	s.App = app.InitTestApp(true)
	s.Ctx = s.App.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: ChainID})
	return s
}

func GetRandomAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()

	bech32Prefix := "temporal"
	accountPubKeyPrefix := bech32Prefix + "pub"
	validatorAddressPrefix := bech32Prefix + "valoper"
	validatorPubKeyPrefix := bech32Prefix + "valoperpub"
	consNodeAddressPrefix := bech32Prefix + "valcons"
	consNodePubKeyPrefix := bech32Prefix + "valconspub"

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(bech32Prefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)

	return sdk.AccAddress(pk.Address())
}
