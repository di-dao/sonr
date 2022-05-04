package transmit

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/sonr/cmd/motor-lib/config"
	"github.com/sonr-io/sonr/internal/host"
	device "github.com/sonr-io/sonr/pkg/fs"
	v1 "go.buf.build/grpc/go/sonr-io/core/host/transmit/v1"
	motor "go.buf.build/grpc/go/sonr-io/core/motor/v1"

	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
)

// TransmitProtocol type
type TransmitProtocol struct {
	callback config.CallbackImpl
	node     host.SonrHost
	ctx      context.Context // Context
	current  *v1.Session     // current session
	mode     device.Role
}

// New creates a new TransferProtocol
func New(ctx context.Context, host host.SonrHost, cb config.CallbackImpl, options ...Option) (*TransmitProtocol, error) {
	// create a new transfer protocol
	protocol := &TransmitProtocol{
		ctx:      ctx,
		node:     host,
		callback: cb,
	}
	// Set options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	opts.Apply(protocol)

	// Setup Stream Handlers
	host.SetStreamHandler(FilePID, protocol.onIncomingTransfer)
	logger.Debug("✅  TransmitProtocol is Activated \n")
	return protocol, nil
}

// CurrentSession returns the current session
func (p *TransmitProtocol) CurrentSession() (*v1.Session, error) {
	if p.current != nil {
		return p.current, nil
	}
	return nil, ErrNoSession
}

// Incoming is called by the node to accept an incoming transfer
func (p *TransmitProtocol) Incoming(payload *types.Payload, from *types.Peer) error {
	// Get User Peer
	to, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get User Peer", err)
		return err
	}

	// Create New TransferEntry
	p.current = NewInSession(payload, from, to)
	return nil
}

// Outgoing is called by the node to initiate a transfer
func (p *TransmitProtocol) Outgoing(payload *types.Payload, to *types.Peer) error {
	// Get User Peer
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer", err)
		return err
	}

	// Get Id
	toId, err := Libp2pID(to)
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer ID", err)
		return err
	}

	// Create New TransferEntry
	p.current = NewOutSession(payload, from, to)

	// Send Files
	if p.current.Payload.GetItems()[0].GetMime().Type != types.MIME_TYPE_URL {
		// Create New Stream
		stream, err := p.node.NewStream(p.ctx, toId, FilePID)
		if err != nil {
			logger.Errorf("%s - Failed to Create New Stream", err)
			return err
		}

		// Start Transfer
		p.onOutgoingTransfer(stream)
	}

	return nil
}

// Reset resets the current session
func (p *TransmitProtocol) Reset(event *motor.OnTransmitCompleteResponse) {
	logger.Debug("Resetting TransmitProtocol")
	p.callback.OnComplete(event)
	p.current = nil
}

// onIncomingTransfer incoming transfer handler
func (p *TransmitProtocol) onIncomingTransfer(stream network.Stream) {
	logger.Debug("Received Incoming Transfer")
	// Find Entry in Queue
	entry, err := p.CurrentSession()
	if err != nil {
		logger.Errorf("%s - Failed to find transfer request", err)
		stream.Close()
		return
	}

	// Create New Reader
	event, err := RouteSessionStream(entry, stream, p.callback)
	if err != nil {
		logger.Errorf("%s - Failed to Read From Stream", err)
		stream.Close()
		return
	}
	p.Reset(event)
}

// onOutgoingTransfer is called by onInviteResponse if Validated
func (p *TransmitProtocol) onOutgoingTransfer(stream network.Stream) {
	logger.Debug("Received Accept Decision, Starting Outgoing Transfer")

	// Find Entry in Queues
	entry, err := p.CurrentSession()
	if err != nil {
		logger.Errorf("%s - Failed to find transfer request", err)
		stream.Close()
		return
	}

	// Create New Writer
	event, err := RouteSessionStream(entry, stream, p.callback)
	if err != nil {
		logger.Errorf("%s - Failed to Write To Stream", err)
		stream.Close()
		return
	}
	p.Reset(event)
}

// Libp2pID returns the PeerID based on PublicKey from Profile
func Libp2pID(p *types.Peer) (peer.ID, error) {
	// Check if PublicKey is empty
	if len(p.GetPublicKey()) == 0 {
		return "", errors.New("Peer Public Key is not set.")
	}

	pubKey, err := crypto.UnmarshalPublicKey(p.GetPublicKey())
	if err != nil {
		return "", err
	}

	// Return Peer ID
	id, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	return id, nil
}
