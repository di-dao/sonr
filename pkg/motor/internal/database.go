package internal

import (
	"github.com/di-dao/sonr/internal/models"
	"github.com/di-dao/sonr/internal/fs"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const kVaultDBFileName = "vault.db"

type Database interface {
	GetCredential(id string) (*models.Credential, error)
	GetProfile(id string) (*models.Profile, error)
	GetWallet(id string) (*models.Wallet, error)

	InsertCredentials(credentials ...*models.Credential) error
	InsertProfiles(profiles ...*models.Profile) error
	InsertWallets(wallets ...*models.Wallet) error

	ListCredentials(origin string) ([]*models.Credential, error)
	ListProfiles() ([]*models.Profile, error)
	ListWallets() ([]*models.Wallet, error)
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

type embedDB struct {
	DB *gorm.DB
}

func (db *embedDB) GetCredential(id string) (*models.Credential, error) {
	credential := new(models.Credential)
	db.DB.First(&credential, "id = ?", id)
	return credential, nil
}

func (db *embedDB) GetProfile(id string) (*models.Profile, error) {
	profile := new(models.Profile)
	db.DB.First(&profile, "id = ?", id)
	return profile, nil
}

func (db *embedDB) GetWallet(id string) (*models.Wallet, error) {
	wallet := new(models.Wallet)
	db.DB.First(&wallet, "id = ?", id)
	return wallet, nil
}

func (db *embedDB) InsertCredentials(credentials ...*models.Credential) error {
	for _, c := range credentials {
		db.DB.Create(c)
	}
	return nil
}

func (db *embedDB) InsertProfiles(profiles ...*models.Profile) error {
	for _, p := range profiles {
		db.DB.Create(p)
	}
	return nil
}

func (db *embedDB) InsertWallets(wallets ...*models.Wallet) error {
	for _, w := range wallets {
		db.DB.Create(w)
	}
	return nil
}

func (db *embedDB) ListCredentials(origin string) ([]*models.Credential, error) {
	var credentials []*models.Credential
	db.DB.Find(&credentials, "origin = ?", origin)
	return credentials, nil
}

func (db *embedDB) ListProfiles() ([]*models.Profile, error) {
	var profiles []*models.Profile
	db.DB.Find(&profiles)
	return profiles, nil
}

func (db *embedDB) ListWallets() ([]*models.Wallet, error) {
	var wallets []*models.Wallet
	db.DB.Find(&wallets)
	return wallets, nil
}
