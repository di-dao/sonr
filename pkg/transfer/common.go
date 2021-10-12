package transfer

import (
	"errors"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/state"
)

// Transfer Emission Events
const (
	Event_INVITED   = "transfer-invited"
	Event_RESPONDED = "transfer-responded"
	Event_PROGRESS  = "transfer-progress"
	Event_COMPLETED = "transfer-completed"
	ITEM_INTERVAL   = 25
)

// Transfer Protocol ID's
const (
	RequestPID  protocol.ID = "/transfer/request/0.0.1"
	ResponsePID protocol.ID = "/transfer/response/0.0.1"
	SessionPID  protocol.ID = "/transfer/session/0.0.1"
)

// Error Definitions
var (
	logger             = golog.Child("protocols/transfer")
	ErrTimeout         = errors.New("Session has Timed out")
	ErrParameters      = errors.New("Failed to create new TransferProtocol, invalid parameters")
	ErrInvalidResponse = errors.New("Invalid Transfer InviteResponse provided to TransferProtocol")
	ErrInvalidRequest  = errors.New("Invalid Transfer InviteRequest provided to TransferProtocol")
	ErrFailedEntry     = errors.New("Failed to get Topmost entry from Queue")
	ErrFailedAuth      = errors.New("Failed to Authenticate message")
	ErrEmptyRequests   = errors.New("Empty Request list provided")
	ErrRequestNotFound = errors.New("Request not found in list")
)

// checkParams Checks if Non-nil Parameters were passed
func checkParams(host *host.SNRHost, em *state.Emitter) error {
	if host == nil {
		logger.Error("Host provided is nil", ErrParameters)
		return ErrParameters
	}
	if em == nil {
		logger.Error("Emitter provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.HasRouting()
}

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *api.DecisionEvent {
	return &api.DecisionEvent{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *api.InviteEvent {
	return &api.InviteEvent{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}
