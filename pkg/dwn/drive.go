package motor

import (
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/internal/fs"
	"github.com/di-dao/sonr/pkg/dwn/internal"
)

const kSonrHRP = "idx"

// vfd is the struct implementation of an IPFS file system
type drive struct {
	kss    kss.Set
	vault  internal.Vault
	folder fs.Folder
	addr   string
}

// NewVFS creates a new virtual file system.
func New() (*drive, error) {
	kss, err := mpc.GenerateKss()
	if err != nil {
		return nil, err
	}

	addr, err := bech32.ConvertAndEncode(kSonrHRP, kss.PublicKey().Bytes())
	if err != nil {
		return nil, err
	}

	rootDir, err := fs.NewVaultFolder(addr)
	if err != nil {
		return nil, err
	}
	vlt, err := internal.InitializeVault(rootDir, kss, addr)
	if err != nil {
		return nil, err
	}

	return &drive{
		folder: rootDir,
		addr:   addr,
		kss:    kss,
		vault:  vlt,
	}, nil
}
