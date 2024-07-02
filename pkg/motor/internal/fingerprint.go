package internal

import (
	"os"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/secret"
	"github.com/di-dao/sonr/internal/fs"
)

func CreateFingerprint(dir fs.Folder, db Database, publicKey crypto.PublicKey) error {
	pk, err := secret.NewKey("credentials", publicKey)
	if err != nil {
		return err
	}
	creds, err := db.ListCredentials()
	if err != nil {
		return err
	}

	credIDStrs := make([]string, len(creds))
	for i, c := range creds {
		credIDStrs[i] = c.DID
	}

	acc, err := pk.CreateAccumulator(credIDStrs...)
	if err != nil {
		return err
	}

	bz, err := acc.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = dir.WriteFile("fingerprint", bz, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
