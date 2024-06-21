package db

import (
	"github.com/di-dao/sonr/internal/vfd/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type VaultDB struct {
	*gorm.DB
}

func SeedTables() (*VaultDB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Account{}, &models.Credential{}, &models.Profile{}, &models.Property{})
	return &VaultDB{DB: db}, nil
}
