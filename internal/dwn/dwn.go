package main

import (
	"github.com/di-dao/sonr/crypto/kss"
	fs "github.com/di-dao/sonr/internal/vfs"
)

type Vault interface {
	ControllerAPI
	Database
}

type vault struct {
	ControllerAPI
	Database
}

func InitializeVault(rootDir fs.Folder, kss kss.Set, addr string) (Vault, error) {
	controller, err := storeKeyshares(rootDir, kss)
	if err != nil {
		return nil, err
	}
	database, err := seedTables(rootDir)
	if err != nil {
		return nil, err
	}
	return &vault{ControllerAPI: controller, Database: database}, nil
}
