package fs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/core/coreiface/options"
	"gorm.io/gorm"
)

// Constant for the name of the folder where the vaults are stored
const kVaultsFolderName = ".sonr-vaults"

// VaultsFolder is the folder where the vaults are stored
var VaultsFolder Folder

type ipfsDB struct {
	DB *gorm.DB
}

// Package initializes the VaultsFolder
func init() {
	// Initialize VaultsFolder
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Create the folder if it does not exist
	VaultsFolder = NewFolder(filepath.Join(homeDir, kVaultsFolderName))
	if !VaultsFolder.Exists() {
		if err := VaultsFolder.Create(); err != nil {
			panic(err)
		}
	}
}

// NewVaultFolder creates a new folder under the VaultsFolder directory
func NewVaultFolder(name string) (Folder, error) {
	vaultFolder := VaultsFolder.Join(name)
	err := vaultFolder.Create()
	if err != nil {
		return "", err
	}
	return vaultFolder, nil
}

// SaveToIPFS saves the Folder to IPFS and returns the IPFS path
func SyncFolderToIPFS(ctx context.Context, f Folder) (path.Path, error) {
	node, err := f.Node()
	if err != nil {
		return nil, err
	}
	c, err := getIPFSClient()
	if err != nil {
		return nil, err
	}
	return c.Unixfs().Add(ctx, node)
}

// PublishToIPNS publishes the Folder to IPNS
func (f Folder) PublishToIPNS(ctx context.Context, ipfsPath path.Path) error {
	c, err := getIPFSClient()
	if err != nil {
		return err
	}
	_, err = c.Name().Publish(ctx, ipfsPath, options.Name.Key(f.Name()))
	return err
}

// LoadFromIPFS loads a Folder from IPFS
func LoadFromIPFS(ctx context.Context, path string) (Folder, error) {
	c, err := getIPFSClient()
	if err != nil {
		return "", err
	}
	cid, err := ParsePath(path)
	if err != nil {
		return "", err
	}
	node, err := c.Unixfs().Get(ctx, cid)
	if err != nil {
		return "", err
	}
	return LoadNodeInFolder(path, node)
}

// LoadNodeInFolder loads an IPFS node into a Folder
func LoadNodeInFolder(path string, node files.Node) (Folder, error) {
	folder := NewFolder(path)
	if err := folder.Create(); err != nil {
		return "", err
	}

	switch n := node.(type) {
	case files.File:
		f, err := os.Create(folder.Path())
		if err != nil {
			return "", err
		}
		defer f.Close()
		if _, err := io.Copy(f, n); err != nil {
			return "", err
		}
	case files.Directory:
		entries := n.Entries()
		for entries.Next() {
			name := entries.Name()
			childNode := entries.Node()
			childPath := filepath.Join(folder.Path(), name)
			if _, err := LoadNodeInFolder(childPath, childNode); err != nil {
				return "", err
			}
		}
		if entries.Err() != nil {
			return "", entries.Err()
		}
	default:
		return "", fmt.Errorf("unsupported node type: %T", n)
	}

	return folder, nil
}
