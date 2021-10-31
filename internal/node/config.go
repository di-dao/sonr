package node

import (
	"context"
	"errors"

	"github.com/kataras/golog"
	api "github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

// Error Definitions
var (
	logger             = golog.Default.Child("internal/node")
	ErrEmptyQueue      = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery    = errors.New("No SName or PeerID provided.")
	ErrMissingParam    = errors.New("Paramater is missing.")
	ErrProtocolsNotSet = errors.New("Node Protocol has not been initialized.")
)

// Option is a function that modifies the node options.
type Option func(*options)

// WithRequest sets the initialize request.
func WithRequest(req *api.InitializeRequest) Option {
	return func(o *options) {
		o.location = req.GetLocation()
		o.profile = req.GetProfile()
		o.connection = req.GetConnection()
	}
}

// WithMode starts the Client RPC server as a highway node.
func WithMode(m StubMode) Option {
	return func(o *options) {
		o.mode = m
	}
}

// options is a collection of options for the node.
type options struct {
	connection common.Connection
	location   *common.Location
	mode       StubMode
	profile    *common.Profile
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *options {
	return &options{
		mode:       StubMode_LIB,
		location:   api.DefaultLocation(),
		connection: common.Connection_WIFI,
		profile:    common.NewDefaultProfile(),
	}
}

// Apply applies Options to node
func (opts *options) Apply(ctx context.Context, node *Node) error {
	// Set Mode
	node.mode = opts.mode

	// Handle by Node Mode
	if opts.mode.HasMotor() {
		logger.Debug("Starting Client stub...")
		// Client Node Type
		stub, err := node.startMotorStub(ctx, opts)
		if err != nil {
			logger.Errorf("%s - Failed to start Client Service", err)
			return err
		}

		// Set Stub to node
		node.motor = stub

	} else {
		logger.Debug("Starting Highway stub...")
		// Highway Node Type
		stub, err := node.startHighwayStub(ctx, opts)
		if err != nil {
			logger.Errorf("%s - Failed to start Highway Service", err)
			return err
		}

		// Set Stub to node
		node.highway = stub
	}
	return nil
}
