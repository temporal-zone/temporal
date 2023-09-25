package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()

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

	return sdk.AccAddress(addr).String()
}
