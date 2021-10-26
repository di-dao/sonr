package transmit

import (
	"container/list"
	"context"
	"fmt"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
)

// TransmitProtocol type
type TransmitProtocol struct {
	node         api.NodeImpl
	ctx          context.Context // Context
	host         *host.SNRHost   // local host
	sessionQueue *SessionQueue   // transfer session queue
	supplyQueue  *list.List      // supply queue
}

// NewProtocol creates a new TransferProtocol
func NewProtocol(ctx context.Context, host *host.SNRHost, node api.NodeImpl) (*TransmitProtocol, error) {
	// Check parameters
	if err := checkParams(host); err != nil {
		logger.Errorf("%s - Failed to create TransmitProtocol", err)
		return nil, err
	}

	// create a new transfer protocol
	invProtocol := &TransmitProtocol{
		ctx:  ctx,
		host: host,
		sessionQueue: &SessionQueue{
			ctx:   ctx,
			host:  host,
			queue: list.New(),
		},
		supplyQueue: list.New(),
		node:        node,
	}

	// Setup Stream Handlers
	host.SetStreamHandler(RequestPID, invProtocol.onInviteRequest)
	host.SetStreamHandler(ResponsePID, invProtocol.onInviteResponse)
	host.SetStreamHandler(SessionPID, invProtocol.onIncomingTransfer)
	logger.Debug("✅  TransmitProtocol is Activated \n")
	return invProtocol, nil
}

// Request Method sends a request to Transfer Data to a remote peer
func (p *TransmitProtocol) Request(to *common.Peer) error {
	// Create Request
	id, req, err := p.createRequest(to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// Check if the response is valid
	if req == nil {
		return ErrInvalidRequest
	}

	// sign the data
	signature, err := p.host.SignMessage(req)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.host.SendMessage(id, RequestPID, req)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}

	// store the request in the map
	return p.sessionQueue.AddOutgoing(id, req)
}

// Respond Method authenticates or declines a Transfer Request
func (p *TransmitProtocol) Respond(decs bool, to *common.Peer) error {
	// Create Response
	id, resp, err := p.createResponse(decs, to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// Check if the response is valid
	if resp == nil {
		return ErrInvalidResponse
	}

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// Send Response
	err = p.host.SendMessage(id, ResponsePID, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}
	return nil
}

// Supply a transfer item to the queue
func (p *TransmitProtocol) Supply(req *api.SupplyRequest) error {
	// Profile from NodeImpl
	profile, err := p.node.Profile()
	if err != nil {
		logger.Errorf("%s - Failed to Get Profile from Node")
		return err
	}

	// Create Transfer
	payload, err := req.ToPayload(profile)
	if err != nil {
		logger.Errorf("%s - Failed to Supply Paths", err)
		return err
	}

	// Add items to transfer
	p.supplyQueue.PushBack(payload)
	logger.Debug(fmt.Sprintf("Added %v items to supply queue.", req.Count()), golog.Fields{"File Count": payload.FileCount(), "URL Count": payload.URLCount()})
	
	// Check if Peer is provided
	if req.GetIsPeerSupply() {
		logger.Debug("Peer Supply Request. Sending Invite after supply")
		err = p.Request(req.GetPeer())
		if err != nil {
			logger.Errorf("%s - Failed to Send Request to Peer", err)
			return err
		}
	}
	return nil
}
