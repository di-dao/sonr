package models

import (
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/ipfs/boxo/files"
)

const kSonrHRP = "idx"

type KeysetFile struct {
	mpcKeys kss.Set
	File    files.File
	Name    string
}

func NewKeysetFile() (*KeysetFile, error) {
	set, err := mpc.GenerateKss()
	if err != nil {
		return nil, err
	}
	addr := formatSonrAddr(set.PublicKey())
	return &KeysetFile{
		File:    nil,
		mpcKeys: set,
		Name:    addr,
	}, nil
}

// utility function for encoding publicKey as standard Sonr Address
func formatSonrAddr(pubKey crypto.PublicKey) string {
	addr, _ := bech32.ConvertAndEncode(kSonrHRP, pubKey.Bytes())
	return addr
}
