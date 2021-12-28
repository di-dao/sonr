package node

import common "github.com/sonr-io/core/common"

// Option is a function that modifies the node options.
type Option func(*options)

// WithMode starts the Client RPC server as a highway node.
func WithMode(m StubMode) Option {
	return func(o *options) {
		o.mode = m
	}
}

// options is a collection of options for the node.
type options struct {
	connection    common.Connection
	location      *common.Location
	mode          StubMode
	profile       *common.Profile
	configuration *Configuration
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *options {
	return &options{
		mode:       StubMode_LIB,
		connection: common.Connection_WIFI,
		profile:    common.NewDefaultProfile(),
	}
}
