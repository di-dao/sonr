package state

import "errors"

// StateType represents an extensible state type in the state machine.
type StateType string

// EventType represents an extensible event type in the state machine.
type EventType string

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}

// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")

const (
	// ### - Default Types -
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"

	// ### - Invite Method Types -
	// 1. States
	// UnknownPeer is state for when Peer is not Found
	UnknownPeer StateType = "UnknownPeer"

	// Pending is state for when Invite Succeeds
	Pending StateType = "Pending"

	// 2. Events
	// FindPeer Method Events
	FailFindPeer    EventType = "FailFindPeer"
	SucceedFindPeer EventType = "SucceedFindPeer"

	// SendInvite Method Events
	FailSendInvite    EventType = "FailSendInvite"
	SucceedSendInvite EventType = "SucceedSendInvite"

	// ### - Respond Method Types -

)
