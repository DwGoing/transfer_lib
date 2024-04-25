package tron

import (
	"transfer_lib/pkg/common"

	"github.com/btcsuite/btcd/btcec"
)

func GetPrivateKeyFromSeed(seed []byte, index int64) (*btcec.PrivateKey, error) {
	return common.GetPrivateKeyFromSeed(seed, common.ExtendedKeyVersion_xprv, "m/44'/195'/0'/0/", index)
}

func NewAccountFromPrivateKey(privateKeyHex *btcec.PrivateKey) (*common.Account, error) {
	return common.NewAccountFromPrivateKey(common.AddressType_TRON, privateKeyHex)
}

func NewAccountFromPrivateKeyHex(privateKeyHex string) (*common.Account, error) {
	return common.NewAccountFromPrivateKeyHex(common.AddressType_TRON, privateKeyHex)
}
