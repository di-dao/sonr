package internal

import (
	"os"

	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/pkg/fs"
)

const (
	kValKSFileName  = "._.val.snr"
	kUserKSFileName = "._.usr.snr"
)

type Controller interface{}

type enclave struct {
	UserKS fs.File
	ValKS  fs.File
}

func StoreKeyshares(folder fs.Folder, set kss.Set) (Controller, error) {
	usrBz, err := set.Usr().Marshal()
	if err != nil {
		return nil, err
	}
	valBz, err := set.Val().Marshal()
	if err != nil {
		return nil, err
	}
	folder.WriteFile(kUserKSFileName, usrBz, os.ModePerm)
	return nil, nil
}
