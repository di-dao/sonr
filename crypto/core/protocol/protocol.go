//
// Copyright Coinbase, Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package protocol

import (
	"fmt"
)

var (
	ErrNotInitialized   = fmt.Errorf("object has not been initialized")
	ErrProtocolFinished = fmt.Errorf("the protocol has finished")
)

const (
	// Dkls18Dkg specifies the DKG protocol of the DKLs18 potocol.
	Dkls18Dkg = "DKLs18-DKG"

	// Dkls18Sign specifies the DKG protocol of the DKLs18 potocol.
	Dkls18Sign = "DKLs18-Sign"

	// Dkls18Refresh specifies the DKG protocol of the DKLs18 potocol.
	Dkls18Refresh = "DKLs18-Refresh"

	// versions will increment in 100 intervals, to leave room for adding other versions in between them if it is
	// ever needed in the future.

	// Version0 is version 0!
	Version0 = 100

	// Version1 is version 2!
	Version1 = 200
)

// Message provides serializers and deserializer for the inputs and outputs of each step of the protocol.
// Moreover, it adds some metadata and versioning around the serialized data.
type Message struct {
	Payloads map[string][]byte
	Metadata map[string]string
	Protocol string
	Version  uint
}

// Iterator an interface for the DKLs18 protocols that follows the iterator pattern.
type Iterator interface {
	// Next runs the next round of the protocol.
	// Returns `ErrProtocolFinished` when protocol has completed.
	Next(input *Message) (*Message, error)

	// Result returns the final result, if any, of the completed protocol.
	// Returns nil if the protocol has not yet terminated.
	// Returns an error if an error was encountered during protocol execution.
	Result(version uint) (*Message, error)
}
