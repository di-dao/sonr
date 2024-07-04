package dwn

import (
	"fmt"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/secret"
	"github.com/di-dao/sonr/internal/orm"
	fs "github.com/di-dao/sonr/internal/vfs"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const kVaultDBFileName = "vault.db"

type Database interface {
	ExistsCredential(did string) bool
	ExistsProfile(did string) bool
	ExistsWallet(did string) bool

	GetCredential(did string) (*orm.Credential, error)
	GetProfile(did string) (*orm.Profile, error)
	GetWallet(did string) (*orm.Wallet, error)

	InsertCredentials(credentials ...*orm.Credential) error
	InsertProfiles(profiles ...*orm.Profile) error
	InsertWallets(wallets ...*orm.Wallet) error

	ListCredentials() ([]*orm.Credential, error)
	ListProfiles() ([]*orm.Profile, error)
	ListWallets() ([]*orm.Wallet, error)

	WitnessCredential(publicKey crypto.PublicKey, did string) ([]byte, error)
	WitnessProfile(publicKey crypto.PublicKey, did string) ([]byte, error)
	WitnessWallet(publicKey crypto.PublicKey, did string) ([]byte, error)
}

func seedTables(dir fs.Folder) (Database, error) {
	file, err := dir.Touch(kVaultDBFileName)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(file.Path()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&orm.Wallet{}, &orm.Credential{}, &orm.Profile{})
	if err != nil {
		return nil, err
	}
	return &embedDB{DB: db}, nil
}

type embedDB struct {
	DB *gorm.DB
}

func (db *embedDB) GetCredential(id string) (*orm.Credential, error) {
	credential := new(orm.Credential)
	db.DB.First(&credential, "did = ?", id)
	return credential, nil
}

func (db *embedDB) GetProfile(id string) (*orm.Profile, error) {
	profile := new(orm.Profile)
	db.DB.First(&profile, "did = ?", id)
	return profile, nil
}

func (db *embedDB) GetWallet(id string) (*orm.Wallet, error) {
	wallet := new(orm.Wallet)
	db.DB.First(&wallet, "did = ?", id)
	return wallet, nil
}

func (db *embedDB) ExistsCredential(did string) bool {
	var count int64
	db.DB.Model(&orm.Credential{}).Where("did = ?", did).Count(&count)
	return count > 0
}

func (db *embedDB) ExistsProfile(did string) bool {
	var count int64
	db.DB.Model(&orm.Profile{}).Where("did = ?", did).Count(&count)
	return count > 0
}

func (db *embedDB) ExistsWallet(did string) bool {
	var count int64
	db.DB.Model(&orm.Wallet{}).Where("did = ?", did).Count(&count)
	return count > 0
}

func (db *embedDB) InsertCredentials(credentials ...*orm.Credential) error {
	for _, c := range credentials {
		db.DB.Create(c)
	}
	return nil
}

func (db *embedDB) InsertProfiles(profiles ...*orm.Profile) error {
	for _, p := range profiles {
		db.DB.Create(p)
	}
	return nil
}

func (db *embedDB) InsertWallets(wallets ...*orm.Wallet) error {
	for _, w := range wallets {
		db.DB.Create(w)
	}
	return nil
}

func (db *embedDB) ListCredentials() ([]*orm.Credential, error) {
	var credentials []*orm.Credential
	db.DB.Find(&credentials)
	return credentials, nil
}

func (db *embedDB) ListProfiles() ([]*orm.Profile, error) {
	var profiles []*orm.Profile
	db.DB.Find(&profiles)
	return profiles, nil
}

func (db *embedDB) ListWallets() ([]*orm.Wallet, error) {
	var wallets []*orm.Wallet
	db.DB.Find(&wallets)
	return wallets, nil
}

func (db *embedDB) WitnessCredential(publicKey crypto.PublicKey, did string) ([]byte, error) {
	// Verify that did exists in the table
	if !db.ExistsCredential(did) {
		return nil, fmt.Errorf("credential with DID %s does not exist", did)
	}

	pk, err := secret.NewKey("credentials", publicKey)
	if err != nil {
		return nil, err
	}

	creds, err := db.ListCredentials()
	if err != nil {
		return nil, err
	}

	credIDStrs := make([]string, len(creds))
	for i, c := range creds {
		credIDStrs[i] = c.DID
	}

	acc, err := pk.CreateAccumulator(credIDStrs...)
	if err != nil {
		return nil, err
	}

	witness, err := pk.CreateWitness(acc, did)
	if err != nil {
		return nil, err
	}
	return witness.MarshalBinary()
}

func (db *embedDB) WitnessProfile(publicKey crypto.PublicKey, did string) ([]byte, error) {
	// Verify that did exists in the table
	if !db.ExistsProfile(did) {
		return nil, fmt.Errorf("profile with DID %s does not exist", did)
	}

	pk, err := secret.NewKey("profiles", publicKey)
	if err != nil {
		return nil, err
	}

	profiles, err := db.ListProfiles()
	if err != nil {
		return nil, err
	}

	profileIDs := make([]string, len(profiles))
	for i, p := range profiles {
		profileIDs[i] = p.DID
	}

	acc, err := pk.CreateAccumulator(profileIDs...)
	if err != nil {
		return nil, err
	}

	witness, err := pk.CreateWitness(acc, did)
	if err != nil {
		return nil, err
	}
	return witness.MarshalBinary()
}

func (db *embedDB) WitnessWallet(publicKey crypto.PublicKey, did string) ([]byte, error) {
	// Verify that did exists in the table
	if !db.ExistsWallet(did) {
		return nil, fmt.Errorf("wallet with DID %s does not exist", did)
	}

	pk, err := secret.NewKey("wallets", publicKey)
	if err != nil {
		return nil, err
	}

	wallets, err := db.ListWallets()
	if err != nil {
		return nil, err
	}

	walletIDs := make([]string, len(wallets))
	for i, w := range wallets {
		walletIDs[i] = w.DID
	}

	acc, err := pk.CreateAccumulator(walletIDs...)
	if err != nil {
		return nil, err
	}

	witness, err := pk.CreateWitness(acc, did)
	if err != nil {
		return nil, err
	}
	return witness.MarshalBinary()
}
