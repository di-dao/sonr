package internal

import (
	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/internal/fs"
	"github.com/di-dao/sonr/internal/models"
)

type FingerprintAPI interface {
	lock(credential *models.Credential)
	unlock(credential *models.Credential)

	Lock(did string)
	Unlock(did string)
}

func Lock(dir fs.Folder) {
}

func Unlock(dir fs.Folder) {
}

type fingerprint struct {
	publicKey crypto.PublicKey
}
