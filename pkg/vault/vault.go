package vault

import (
	"context"

	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/internal/local"
	"github.com/di-dao/sonr/pkg/ipfs"
)

// Vault is an interface that defines the methods for a vault.
type Vault interface{}

// vault is a struct that contains the information of a vault to be stored in the vault
type vault struct {
	vltFS *vaultFS
	vfs   ipfs.VFS
}

// New creates a new vault from a set of keyshares.
func Generate(ctx context.Context) (Vault, error) {
	snrCtx := local.UnwrapContext(ctx)
	// Generate keyshare
	keyshares, err := mpc.GenerateKss()
	if err != nil {
		return nil, err
	}
	fs, err := createVaultFS(keyshares)
	if err != nil {
		return nil, err
	}

	// Update the context with the wallet address
	snrCtx.UserAddress = fs.Wallet.SonrAddress()
	local.WrapContext(snrCtx)
	vaultCache.Set(contextKey(snrCtx.SessionID), fs)

	// Create a new vault
	return &vault{
		vltFS: fs,
		vfs:   ipfs.NewFSWithKss(keyshares, fs.Wallet.SonrAddress()),
	}, nil
}

// Connect connects to an existing vault.
func Connect(ctx context.Context, address string) (Vault, error) {
	snrCtx := local.UnwrapContext(ctx)
	vfs, err := ipfs.GetFileSystem(ctx, address)
	if err != nil {
		return nil, err
	}
	fs, err := loadVaultFS(vfs)
	if err != nil {
		return nil, err
	}

	// Update the context with the wallet address
	snrCtx.UserAddress = fs.Wallet.SonrAddress()
	local.WrapContext(snrCtx)

	// Create a new vault
	return &vault{
		vfs:   vfs,
		vltFS: fs,
	}, nil
}
