package node

import (
	context "context"
	"fmt"
	"net"

	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/mailbox"
	"github.com/sonr-io/core/pkg/transfer"
	grpc "google.golang.org/grpc"
)

// StubMode is the type of the node (Client, Highway)
type StubMode int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	StubMode_LIB StubMode = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	StubMode_CLI

	// StubMode_BIN is the Node utilized for Desktop background process
	StubMode_BIN

	// StubMode_HIGHWAY is the Custodian Node that manages Network
	StubMode_HIGHWAY
)

// IsLib returns true if the node is a client node.
func (m StubMode) IsLib() bool {
	return m == StubMode_LIB
}

// IsBin returns true if the node is a bin node.
func (m StubMode) IsBin() bool {
	return m == StubMode_BIN
}

// IsCLI returns true if the node is a CLI node.
func (m StubMode) IsCLI() bool {
	return m == StubMode_CLI
}

// IsHighway returns true if the node is a highway node.
func (m StubMode) IsHighway() bool {
	return m == StubMode_HIGHWAY
}

// HasClient returns true if the node has a client.
func (m StubMode) HasClient() bool {
	return m.IsLib() || m.IsBin() || m.IsCLI()
}

// HasHighway returns true if the node has a highway stub.
func (m StubMode) HasHighway() bool {
	return m.IsHighway()
}

// Prefix returns golog prefix for the node.
func (m StubMode) Prefix() string {
	var name string
	switch m {
	case StubMode_LIB:
		name = "lib"
	case StubMode_CLI:
		name = "cli"
	case StubMode_BIN:
		name = "bin"
	case StubMode_HIGHWAY:
		name = "highway"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
}

// ClientNodeStub is the RPC Service for the Default Node.
type ClientNodeStub struct {
	// Interfaces
	ClientServiceServer

	// Properties
	ctx        context.Context
	listener   net.Listener
	grpcServer *grpc.Server
	node       *Node

	// Protocols
	*transfer.TransferProtocol
	*exchange.ExchangeProtocol
	*lobby.LobbyProtocol
	*mailbox.MailboxProtocol
}

// startClientService creates a new Client service stub for the node.
func (n *Node) startClientService(ctx context.Context, opts *options) (*ClientNodeStub, error) {
	// Set Transfer Protocol
	transferProtocol, err := transfer.NewProtocol(ctx, n.host, n)
	if err != nil {
		logger.Error("Failed to start TransferProtocol", err)
		return nil, err
	}

	// Set Local Lobby Protocol if Location is provided
	lobbyProtocol, err := lobby.NewProtocol(ctx, n.host, n, lobby.WithLocation(opts.location))
	if err != nil {
		logger.Error("Failed to start LobbyProtocol", err)
		return nil, err
	}

	// Set Exchange Protocol
	var exchProtocol *exchange.ExchangeProtocol
	if opts.mode.IsCLI() {
		exchProtocol, err = exchange.NewProtocol(ctx, n.host, n, exchange.TempName(opts.profile.SName))
		if err != nil {
			logger.Error("Failed to start ExchangeProtocol", err)
			return nil, err
		}
	} else {
		exchProtocol, err = exchange.NewProtocol(ctx, n.host, n)
		if err != nil {
			logger.Error("Failed to start ExchangeProtocol", err)
			return nil, err
		}
	}

	// // Set Mailbox Protocol
	// mailboxProtocol, err := mailbox.NewProtocol(ctx, n.host, n.Emitter)
	// if err != nil {
	// 	logger.Error("Failed to start MailboxProtocol", err)
	// 	return nil, err
	// }

	// Open Listener on Port
	listener, err := net.Listen(opts.network, opts.Address())
	if err != nil {
		logger.Fatal("Failed to bind listener to port ", err)
		return nil, err
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	stub := &ClientNodeStub{
		ctx:              ctx,
		TransferProtocol: transferProtocol,
		ExchangeProtocol: exchProtocol,
		LobbyProtocol:    lobbyProtocol,
		//MailboxProtocol:  mailboxProtocol,
		grpcServer: grpcServer,
		node:       n,
		listener:   listener,
	}

	// Start Routines
	RegisterClientServiceServer(grpcServer, stub)
	go stub.Serve(ctx, listener)
	return stub, nil
}

// HasProtocols returns true if the node has the protocols.
func (s *ClientNodeStub) HasProtocols() bool {
	return s.TransferProtocol != nil && s.ExchangeProtocol != nil && s.LobbyProtocol != nil
}

// Close closes the RPC Service.
func (s *ClientNodeStub) Close() error {
	s.LobbyProtocol.Close()
	s.ExchangeProtocol.Close()
	s.grpcServer.Stop()
	s.listener.Close()
	return nil
}

// Serve serves the RPC Service on the given port.
func (s *ClientNodeStub) Serve(ctx context.Context, listener net.Listener) {
	// Handle Node Events
	if err := s.grpcServer.Serve(listener); err != nil {
		logger.Error("Failed to serve gRPC", err)
	}
	for {
		// Stop Serving if context is done
		select {
		case <-ctx.Done():
			return
		}
	}
}

// Update method updates the node's properties in the Key/Value Store and Lobby
func (s *ClientNodeStub) Update() error {
	// Call Internal Edit
	peer, err := s.node.Peer()
	if err != nil {
		logger.Error("Failed to get Peer Ref", err)
		return err
	}

	// Check for Valid Protocols
	if s.HasProtocols() {
		// Update LobbyProtocol
		err = s.LobbyProtocol.Update()
		if err != nil {
			logger.Error("Failed to Update Lobby", err)
		} else {
			logger.Info("🌎 Succesfully updated Lobby.")
		}

		// Update ExchangeProtocol
		err := s.ExchangeProtocol.Put(peer)
		if err != nil {
			logger.Error("Failed to Update Exchange", err)
		} else {
			logger.Info("🌎 Succesfully updated Exchange.")
		}
		return err
	} else {
		return ErrProtocolsNotSet
	}
}

// HighwayNodeStub is the RPC Service for the Custodian Node.
type HighwayNodeStub struct {
	HighwayServiceServer
	ClientServiceServer
	*Node

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
	listener   net.Listener
}

// startHighwayService creates a new Highway service stub for the node.
func (n *Node) startHighwayService(ctx context.Context, opts *options) (*HighwayNodeStub, error) {

	// Create the listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_SERVER_PORT))
	if err != nil {
		return nil, err
	}

	// Create the RPC Service
	grpcServer := grpc.NewServer()
	stub := &HighwayNodeStub{
		Node:       n,
		ctx:        ctx,
		grpcServer: grpcServer,
		listener:   listener,
	}
	// Register the RPC Service
	RegisterHighwayServiceServer(stub.grpcServer, stub)
	go stub.Serve(ctx, listener)
	return stub, nil
}

// Serve serves the RPC Service on the given port.
func (s *HighwayNodeStub) Serve(ctx context.Context, listener net.Listener) {
	// Handle Node Events
	if err := s.grpcServer.Serve(s.listener); err != nil {
		logger.Error("Failed to serve gRPC", err)

	}
	for {
		// Stop Serving if context is done
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (s *HighwayNodeStub) Close() error {
	s.listener.Close()
	s.grpcServer.Stop()
	return nil
}
