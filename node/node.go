package node

import (
	"context"
	"net"
	sync "sync"

	"git.mills.io/prologic/bitcask"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/node/api"
	"github.com/sonr-io/core/node/highway"
	"github.com/sonr-io/core/node/motor"
	"github.com/sonr-io/core/pkg/identity"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Standard Node Implementation
	api.NodeImpl
	motor   *motor.MotorStub
	highway *highway.HighwayStub
	mode    api.StubMode

	// Host and context
	host     *host.SNRHost
	listener net.Listener

	// Properties
	ctx      context.Context
	identity *identity.IdentityProtocol
	store    *bitcask.Bitcask
	state    *api.State
	once     sync.Once
}

// NewNode Creates a node with its implemented protocols
func NewNode(ctx context.Context, l net.Listener, options ...Option) (api.NodeImpl, *api.InitializeResponse, error) {
	// Set Node Options
	opts := defaultNodeOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Initialize Host
	host, err := host.NewHost(ctx, host.WithConnection(opts.connection))
	if err != nil {
		logger.Errorf("%s - Failed to initialize host", err)
		return nil, api.NewInitialzeResponse(nil, false), err
	}

	// Open Store with profileBuf
	// Create Node
	node := &Node{
		ctx:      ctx,
		listener: l,
		host:     host,
	}

	logger.Debugf("Opening Store with profile: %s", opts.profile)
	node.identity, err = identity.New(ctx, host, node, identity.WithProfile(opts.profile))
	if err != nil {
		logger.Errorf("%s - Failed to initialize identity", err)
		return nil, api.NewInitialzeResponse(nil, false), err
	}

	// Initialize Stub
	err = opts.Apply(ctx, host, node)
	if err != nil {
		logger.Errorf("%s - Failed to initialize stub", err)
		return nil, api.NewInitialzeResponse(nil, false), err
	}
	// Begin Background Tasks
	return node, api.NewInitialzeResponse(node.Profile, false), nil
}

// GetState returns the current state of the API
func (n *Node) GetState() *api.State {
	n.once.Do(func() {
		chn := make(chan bool)
		close(chn)
		n.state = &api.State{Chn: chn}
	})
	return n.state
}

// Profile returns the profile for the user from diskDB
func (n *Node) Profile() (*common.Profile, error) {
	return n.identity.Profile()
}

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	return n.identity.Peer()
}

// Close closes the node
func (n *Node) Close() {
	// Close Client Stub
	if n.mode.Motor() {
		if err := n.motor.Close(); err != nil {
			logger.Errorf("%s - Failed to close Client Stub, ", err)
		}
	}

	// Close Highway Stub
	if n.mode.IsFull() {

	}

	// Close Store
	if err := n.store.Close(); err != nil {
		logger.Errorf("%s - Failed to close store, ", err)
	}

	// Close Host
	if err := n.host.Close(); err != nil {
		logger.Errorf("%s - Failed to close host, ", err)
	}
}
