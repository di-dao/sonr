package internal

import (
	"os"

	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/internal/fs"
)

const (
	kValKSFileName  = "._.val.snr"
	kUserKSFileName = "._.usr.snr"
)

type ControllerAPI interface {
	SignMessage(msg []byte) ([]byte, error)
	VerifyMessage(msg []byte, sig []byte) (bool, error)
}

type enclave struct {
	UserKS fs.File
	ValKS  fs.File
}

func (e *enclave) SignMessage(msg []byte) ([]byte, error) {
	set, err := recoverKeyshares(e)
	if err != nil {
		return nil, err
	}
	userSign, err := set.Usr().GetSignFunc(msg)
	if err != nil {
		return nil, err
	}
	valSign, err := set.Val().GetSignFunc(msg)
	if err != nil {
		return nil, err
	}
	return mpc.RunSignProtocol(valSign, userSign)
}

func (e *enclave) VerifyMessage(msg []byte, sig []byte) (bool, error) {
	set, err := recoverKeyshares(e)
	if err != nil {
		return false, err
	}
	return set.PublicKey().VerifySignature(msg, sig), nil
}

func recoverKeyshares(encl *enclave) (kss.Set, error) {
	uzerBz, err := encl.UserKS.Read()
	if err != nil {
		return nil, err
	}
	valBz, err := encl.ValKS.Read()
	if err != nil {
		return nil, err
	}
	return kss.LoadKeyshareSet(valBz, uzerBz)
}

func storeKeyshares(folder fs.Folder, set kss.Set) (ControllerAPI, error) {
	usrBz, err := set.Usr().Marshal()
	if err != nil {
		return nil, err
	}
	valBz, err := set.Val().Marshal()
	if err != nil {
		return nil, err
	}
	usrFile, err := folder.WriteFile(kUserKSFileName, usrBz, os.ModePerm)
	if err != nil {
		return nil, err
	}
	valFile, err := folder.WriteFile(kValKSFileName, valBz, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &enclave{usrFile, valFile}, nil
}
