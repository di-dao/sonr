package internal

import (
	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/internal/fs"
)

func Lock(dir fs.Folder) {
}

func Unlock(dir fs.Folder) {
}

type fingerprint struct {
	PublicKey crypto.PublicKey
	Address   string
}
