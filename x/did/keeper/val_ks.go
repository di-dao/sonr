package keeper

import (
	"golang.org/x/crypto/sha3"

	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1/dkg"
	"github.com/di-dao/core/x/did/types"
)

// ValidatorKSOutput is the protocol result for the validator keyshare output
type ValidatorKSOutput = *dkg.AliceOutput

// ValidatorSignFunc is the type for the validator sign function
type ValidatorSignFunc = *dklsv1.AliceSign

// ValidatorRefreshFunc is the type for the validator refresh function
type ValidatorRefreshFunc = *dklsv1.AliceRefresh

// ValidatorKeyshare is the protocol result for the validator keyshare
type ValidatorKeyshare struct {
	Keyshare
	valKSS *protocol.Message
	pubKey *types.PublicKey
}

// createValidatorKeyshare creates a new ValidatorKeyshare
func createValidatorKeyshare(valKSS *protocol.Message) *ValidatorKeyshare {
	return &ValidatorKeyshare{
		valKSS: valKSS,
	}
}

// GetSignFunc returns the sign function for the validator keyshare
func (v *ValidatorKeyshare) GetSignFunc(msg []byte) (ValidatorSignFunc, error) {
	curve := curves.P256()
	aliceSign, err := dklsv1.NewAliceSign(curve, sha3.New256(), msg, v.valKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return aliceSign, nil
}

// GetRefreshFunc returns the refresh function for the validator keyshare
func (v *ValidatorKeyshare) GetRefreshFunc() (ValidatorRefreshFunc, error) {
	curve := curves.P256()
	aliceRefresh, err := dklsv1.NewAliceRefresh(curve, v.valKSS, protocol.Version1)
	if err != nil {
		return nil, err
	}
	return aliceRefresh, nil
}

// PublicKey is the public key for the keyshare
func (u *ValidatorKeyshare) PublicKey() *types.PublicKey {
	aliceOut, err := dklsv1.DecodeAliceDkgResult(u.valKSS)
	if err != nil {
		panic(err)
	}
	pub := &types.PublicKey{
		Key:     aliceOut.PublicKey.ToAffineUncompressed(),
		KeyType: "ecdsa-secp256k1",
	}
	return pub
}
