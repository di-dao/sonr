package exchange

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	common "github.com/sonr-io/core/common"
	"github.com/sonr-io/core/types/go/node/motor/v1"
)

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *motor.OnTransmitDecisionResponse {
	return &motor.OnTransmitDecisionResponse{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *motor.OnTransmitInviteResponse {
	return &motor.OnTransmitInviteResponse{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}

// createRequest creates a new InviteRequest
func (p *ExchangeProtocol) createRequest(to *common.Peer, payload *common.Payload) (peer.ID, *InviteRequest, error) {
	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Fetch Peer ID from Public Key
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}

	// Create new Metadata
	// meta, err := wallet.CreateMetadata(p.host.ID())
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
	// 	return "", nil, err
	// }

	// Create Invite Request
	req := &InviteRequest{
		Payload: payload,
		// TODO: Implement Signed Meta to Proto Method
		// Metadata: api.SignedMetadataToProto(meta),
		To:   to,
		From: from,
	}
	return toId, req, nil
}

// createResponse creates a new InviteResponse
func (p *ExchangeProtocol) createResponse(decs bool, to *common.Peer) (peer.ID, *InviteResponse, error) {

	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Create new Metadata
	// meta, err := wallet.CreateMetadata(p.host.ID())
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
	// 	return "", nil, err
	// }

	// Create Invite Response
	resp := &InviteResponse{
		Decision: decs,
		// TODO: Implement Signed Meta to Proto Method
		//Metadata: api.SignedMetadataToProto(meta),
		From: from,
		To:   to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
}
