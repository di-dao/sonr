package models

import (
	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/bip32"
	"gorm.io/gorm"
)

// Account is a struct that contains the information of a wallet account
type Account struct {
	gorm.Model
	Address    string `json:"address"`
	Controller string `json:"controller"`
	PublicKey  []byte `json:"publicKey"`
	Index      int    `json:"index"`
	CoinType   int64  `json:"coinType"`
}

// TODO - Convert to BeforeCreate cmd
// NewAccount creates a new account from a public key, coin, and index
func NewAccount(pubkey crypto.PublicKey, coin Coin, index int) (*Account, error) {
	expbz := pubkey.Bytes()
	pubBz, err := bip32.ComputePublicKey(expbz, coin.GetPath(), index)
	if err != nil {
		return nil, err
	}
	addr, err := coin.FormatAddress(pubBz)
	if err != nil {
		return nil, err
	}
	return &Account{
		PublicKey: pubBz,
		Index:     index,
		Address:   addr,
		CoinType:  coin.GetIndex(),
	}, nil
}
