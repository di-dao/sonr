package node

import (
	"errors"
	"strings"

	"github.com/libp2p/go-libp2p-core/crypto"
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
)

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	// Get Public Key
	pubKey, err := device.KeyChain.GetPubKey(device.Account)
	if err != nil {
		return nil, err
	}

	// Find PublicKey Buffer
	deviceStat := device.Stat()
	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		logger.Error("Failed to marshal public key", zap.Error(err))
		return nil, err
	}

	// Return Peer
	return &common.Peer{
		SName:     strings.ToLower(n.profile.SName),
		Status:    common.Peer_ONLINE,
		Profile:   n.profile,
		PublicKey: pubBuf,
		Device: &common.Peer_Device{
			HostName: deviceStat.HostName,
			Os:       deviceStat.Os,
			Id:       deviceStat.Id,
			Arch:     deviceStat.Arch,
		},
	}, nil
}

// ToExchangeQueryRequest converts a query request to an exchange query request.
func (f *FindRequest) ToExchangeQueryRequest() (*exchange.QueryRequest, error) {
	if f.GetSName() != "" {
		return &exchange.QueryRequest{
			SName: f.GetSName(),
		}, nil
	}

	if f.GetPeerId() != "" {
		return &exchange.QueryRequest{
			PeerId: f.GetPeerId(),
		}, nil
	}
	return nil, errors.New("No SName or PeerID provided.")
}

// ToFindResponse converts PeerInfo to a FindResponse.
func ToFindResponse(p *common.PeerInfo) *FindResponse {
	return &FindResponse{
		Success: true,
		Peer:    p.Peer,
		PeerId:  p.PeerID.String(),
		SName:   p.SName,
	}
}
