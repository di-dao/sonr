package internal

import (
	"os"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/accumulator"
	"github.com/di-dao/sonr/crypto/secret"
	fs "github.com/di-dao/sonr/internal/vfs"
)

// CreateFingerprint creates a fingerprint for the given database and public key
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

// ValidateAndPurgeFingerprint validates the fingerprint and purges it if it's valid
func ValidateAndPurgeFingerprint(dir fs.Folder, witness []byte, publicKey crypto.PublicKey) (bool, error) {
	pk, err := secret.NewKey("credentials", publicKey)
	if err != nil {
		return false, err
	}
	creds := new(accumulator.Accumulator)
	membership := new(accumulator.MembershipWitness)
	err = membership.UnmarshalBinary(witness)
	if err != nil {
		return false, err
	}

	bz, err := dir.ReadFile("fingerprint")
	if err != nil {
		return false, err
	}

	err = creds.UnmarshalBinary(bz)
	if err != nil {
		return false, err
	}

	if err := pk.VerifyWitness(creds, membership); err != nil {
		return false, err
	}

	err = dir.DeleteFile("fingerprint")
	if err != nil {
		return false, err
	}
	return true, nil
}
