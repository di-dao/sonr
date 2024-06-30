package internal

import (
	"github.com/di-dao/sonr/internal/models"
	"github.com/di-dao/sonr/pkg/fs"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const kVaultDBFileName = "vault.db"

type Database interface{}

type embedDB struct {
	DB *gorm.DB
}

func SeedTables(dir fs.Folder) (Database, error) {
	file, err := dir.Touch(kVaultDBFileName)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(file.Path()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Wallet{}, &models.Credential{}, &models.Profile{})
	if err != nil {
		return nil, err
	}
	return &embedDB{DB: db}, nil
}
