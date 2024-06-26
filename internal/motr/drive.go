package motr

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/crypto/mpc"
	"github.com/di-dao/sonr/pkg/fs"
	"github.com/ipfs/kubo/core/coreiface"
)

const kSonrHRP = "idx"

// VFD is an interface for interacting with a virtual file drive.
type VFD interface {
	Name() string
	GenerateKey(ctx context.Context) (coreiface.Key, error)
}

// vfd is the struct implementation of an IPFS file system
type drive struct {
	folder fs.Folder
	addr   string
	kss    kss.Set
	db     fs.File
	meta   fs.File
}

// NewVFS creates a new virtual file system.
func New() (VFD, error) {
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

	return &drive{
		folder: rootDir,
		addr:   addr,
		kss:    kss,
	}, nil
}

// Name returns the name of the virtual file system.
func (v *drive) Name() string {
	return v.folder.Name()
}

// GenerateKey creates a new IPFS key.
func (v *drive) GenerateKey(ctx context.Context) (coreiface.Key, error) {
	// Implement the key generation logic here
	// You may need to use the IPFS client to generate the key
	return nil, nil
}
