package ipfs

import (
	"context"

	"github.com/di-dao/sonr/internal/local"
	"github.com/di-dao/sonr/pkg/fs"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	coreiface "github.com/ipfs/kubo/core/coreiface"
	"github.com/ipfs/kubo/core/coreiface/options"
)

type IPNSKey = coreiface.Key

// IPFSClient is an interface for interacting with an IPFS node.
type IPFSClient struct {
	*rpc.HttpApi
}

// NewKey creates a new IPFS key.
func NewKey(ctx context.Context, addr string) (IPNSKey, error) {
	// Call the IPFS client to create a new key
	c, err := local.GetIPFSClient()
	if err != nil {
		return nil, err
	}
	key, err := c.Key().Generate(ctx, addr)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// GetFileSystem returns the Folder interface from the client UnixFS API.
func GetFileSystem(ctx context.Context, path string) (*fs.Folder, error) {
	c, err := local.GetIPFSClient()
	if err != nil {
		return nil, err
	}
	cid, err := parsePath(path)
	if err != nil {
		return nil, err
	}

	// Call the IPFS client to get the UnixFS API.
	node, err := c.Unixfs().Get(context.Background(), cid)
	if err != nil {
		return nil, err
	}
	return fs.NewFolder(path, node)
}

// SaveFileSystem saves the Folder interface to the client UnixFS API.
func SaveFileSystem(ctx context.Context, folder *fs.Folder) error {
	// Call the IPFS client to get the UnixFS API.
	c, err := local.GetIPFSClient()
	if err != nil {
		return err
	}
	node, err := folder.Node()
	if err != nil {
		return err
	}
	api, err := c.Unixfs().Add(ctx, node)
	if err != nil {
		return err
	}
	folder.SetPath(api.String())
	return nil
}

// PublishFileSystem publishes the Folder interface to the client UnixFS API.
func PublishFileSystem(ctx context.Context, folder *fs.Folder) error {
	// Call the IPFS client to get the UnixFS API.
	c, err := local.GetIPFSClient()
	if err != nil {
		return err
	}

	err = folder.Validate()
	if err != nil {
		return err
	}

	_, err = c.Name().Publish(ctx, path.New(folder.Path()), options.Name.Key(folder.Name()))
	if err != nil {
		return err
	}
	return nil
}

func parsePath(p string) (path.Path, error) {
	return path.NewPath(p)
}
