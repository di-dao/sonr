package vault

import (
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/pkg/auth"
	"github.com/di-dao/sonr/pkg/ipfs"
	"github.com/di-dao/sonr/pkg/vault/chain"
	"github.com/di-dao/sonr/pkg/vault/controller"
	"github.com/di-dao/sonr/pkg/vault/props"
	"github.com/di-dao/sonr/pkg/vault/wallet"
)

// Vault is an interface that defines the methods for a vault.
type Vault interface {
	AddCredential(origin string, credential *auth.Credential)
	ListCredentials(origin string) []*auth.Credential
	RemoveCredential(origin string, credential *auth.Credential)
}

// vault is a struct that contains the information of a vault to be stored in the vault
type vault struct {
	vfs        ipfs.VFS
	wallet     *wallet.Wallet
	controller controller.Controller
	properties props.Properties
}

// New creates a new vault from a set of keyshares.
func New(keyshares kss.Set) (Vault, error) {
	// Get sonr address and bitcoin address from keyshares
	wallet, err := wallet.New(keyshares)
	if err != nil {
		return nil, err
	}
	return &vault{
		wallet:     wallet,
		properties: props.NewProperties(),
		controller: controller.New(keyshares),
		vfs:        ipfs.NewVFS(wallet.Accounts[chain.CoinSNRType][0].Address),
	}, nil
}
