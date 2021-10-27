package transmit

import (
	"container/list"
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
)

// Session is a single entry in the Transmit queue.
type Session struct {
	direction   common.Direction
	from        *common.Peer
	to          *common.Peer
	payload     *common.Payload
	lastUpdated int64
}

// IsIncoming returns true if the session is incoming.
func (s Session) IsIncoming() bool {
	return s.direction == common.Direction_INCOMING
}

// IsOutgoing returns true if the session is outgoing.
func (s Session) IsOutgoing() bool {
	return s.direction == common.Direction_OUTGOING
}

// IsOutgoing returns true if the session is outgoing.
func (s Session) ReplaceItem(idx int, pi *common.Payload_Item) {
	// Replace Item
	s.payload.Items[idx] = pi

	// Update Last Updated
	s.lastUpdated = int64(time.Now().Unix())
}

// ReadFrom reads the next Session from the given stream.
func (s Session) ReadFrom(stream network.Stream, n api.NodeImpl) (*api.CompleteEvent, error) {
	// Initialize Params
	logger.Debug("Beginning INCOMING Transmit Stream")

	// Handle incoming stream
	rs := msgio.NewReader(stream)
	var wg sync.WaitGroup

	// Write All Files
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)

		// Configure Reader
		config := itemConfig{
			index:  i,
			count:  s.Count(),
			item:   v,
			node:   n,
			reader: rs,
			wg:     wg,
		}

		// Create Reader
		pi, err := handleItemRead(config)
		if err != nil {
			logger.Errorf("%s - Failed to create new reader.", err)
			rs.Close()
			return nil, err
		}
		s.ReplaceItem(i, pi)
	}
	wg.Wait()
	stream.Close()

	// Return Complete Event
	return &api.CompleteEvent{
		From:       s.from,
		To:         s.to,
		Direction:  s.direction,
		Payload:    s.payload,
		CreatedAt:  s.payload.GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
	}, nil
}

// WriteTo writes the Session to the given stream.
func (s Session) WriteTo(stream network.Stream, n api.NodeImpl) (*api.CompleteEvent, error) {
	// Initialize Params
	logger.Debug("Beginning OUTGOING Transmit Stream")
	wc := msgio.NewWriter(stream)
	var wg sync.WaitGroup

	// Create New Writer
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)

		// Configure Writer
		config := itemConfig{
			index:  i,
			count:  s.Count(),
			item:   v,
			node:   n,
			writer: wc,
			wg:     wg,
		}

		// Create New Writer
		err := handleItemWrite(config)
		if err != nil {
			logger.Errorf("%s - Failed to create new writer.", err)
			wc.Close()
			return nil, err
		}

	}

	// Wait for all writes to finish
	wg.Wait()

	// Return Complete Event
	return &api.CompleteEvent{
		From:       s.from,
		To:         s.to,
		Direction:  s.direction,
		Payload:    s.payload,
		CreatedAt:  s.payload.GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
	}, nil
}

// Count returns the number of items in Payload
func (s Session) Count() int {
	return len(s.payload.GetItems())
}

// MapItems performs PayloadItemFunc on each item in the Payload.
func (s Session) Items() []*common.Payload_Item {
	return s.payload.GetItems()
}

// SessionQueue is a queue for incoming and outgoing requests.
type SessionQueue struct {
	ctx   context.Context
	host  *host.SNRHost
	queue *list.List
}

// AddIncoming adds Incoming Request to Transfer Queue
func (sq *SessionQueue) AddIncoming(from peer.ID, req *InviteRequest) error {
	// Authenticate Message
	valid := sq.host.AuthenticateMessage(req, req.Metadata)
	if !valid {
		return ErrFailedAuth
	}

	// Create New TransferEntry
	entry := Session{
		direction:   common.Direction_INCOMING,
		payload:     req.GetPayload(),
		from:        req.GetFrom(),
		to:          req.GetTo(),
		lastUpdated: int64(time.Now().Unix()),
	}

	// Add to Requests
	sq.queue.PushBack(entry)
	return nil
}

// AddOutgoing adds Outgoing Request to Transfer Queue
func (sq *SessionQueue) AddOutgoing(to peer.ID, req *InviteRequest) error {
	// Create New TransferEntry
	entry := Session{
		direction:   common.Direction_OUTGOING,
		payload:     req.GetPayload(),
		from:        req.GetFrom(),
		to:          req.GetTo(),
		lastUpdated: int64(time.Now().Unix()),
	}

	// Add to Requests
	sq.queue.PushBack(entry)
	return nil
}

// Next returns topmost entry in the queue.
func (sq *SessionQueue) Next() (Session, error) {
	// Find Entry for Peer
	entry := sq.queue.Remove(sq.queue.Front()).(Session)
	entry.lastUpdated = int64(time.Now().Unix())
	return entry, nil
}

// Validate takes list of Requests and returns true if Request exists in List and UUID is verified.
// Method also returns the InviteRequest that points to the Response.
func (sq *SessionQueue) Validate(resp *InviteResponse) (Session, error) {
	// Authenticate Message
	valid := sq.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		return Session{}, ErrFailedAuth
	}

	// Check Decision
	if !resp.GetDecision() {
		return Session{}, nil
	}

	// Check if the request is valid
	if sq.queue.Len() == 0 {
		return Session{}, ErrEmptyRequests
	}

	// Get Next Entry
	entry, err := sq.Next()
	if err != nil {
		logger.Errorf("%s - Failed to get Transmit entry", err)
		return Session{}, err
	}

	entry.lastUpdated = int64(time.Now().Unix())
	return entry, nil
}
