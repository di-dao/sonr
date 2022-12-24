package node

import (
	"context"
	"fmt"
	"sync"

	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	mpc "github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// `OnlineNetwork` is a struct that contains a `context.Context`, a `node.Node`, a list of `peer.ID`s, a list
// of `party.ID`s, a `sync.Mutex`, a `chan bool`, a `chan error`, a `chan *mpc.Message`, and a
// `map[party.ID]icore.PubSubSubscription`.
// @property ctx - the context of the network
// @property selfNode - The node that represents the current party.
// @property {[]peer.ID} peerIds - a list of peer IDs that are connected to this node
// @property {[]party.ID} partyIds - a list of party IDs that this node is connected to
// @property mtx - a mutex to protect the network from concurrent access
// @property doneChan - a channel that will be closed when the network is stopped
// @property errorChan - This channel is used to send errors to the caller.
// @property msgInChan - This is the channel that the network will use to send messages to the
// application.
// @property subscriptions - a map of party IDs to PubSub subscriptions.
type OnlineNetwork struct {
	nodes   []*Node
	parties party.IDSlice

	mtx  sync.Mutex
	done chan struct{}

	closedListenChan chan *mpc.Message
	listenChannels   map[party.ID]chan *mpc.Message
	subscriptions    map[party.ID]icore.PubSubSubscription
}

// It creates a new network object, assigns the subscriptions, and returns the network object
func createOnlineNetwork(ctx context.Context, nodes []*Node) (*OnlineNetwork, error) {
	// Convert the peer IDs to party IDs.
	partyIds := make([]party.ID, len(nodes))
	for i, n := range nodes {
		partyIds[i] = n.partyId
	}
	closed := make(chan *protocol.Message)
	close(closed)
	// Create the network object.
	net := &OnlineNetwork{
		nodes:            nodes,
		parties:          partyIds,
		done:             make(chan struct{}),
		listenChannels:   make(map[party.ID]chan *mpc.Message, 2*len(partyIds)),
		closedListenChan: closed,
	}
	return net, nil
}

// Initializing the network.
func (n *OnlineNetwork) init() {
	N := len(n.parties)
	for _, id := range n.parties {
		n.listenChannels[id] = make(chan *protocol.Message, N*N)
	}
	n.done = make(chan struct{})

	for _, node := range n.nodes {
		sub, err := node.PubSub().Subscribe(context.Background(), topicKey(node.partyId))
		if err != nil {
			panic(err)
		}
		n.subscriptions[node.partyId] = sub
		go n.handleSubscription(node.partyId, sub)
	}
}

// Returning the channel that the network will use to send messages to the application.
func (n *OnlineNetwork) Next(id party.ID) <-chan *mpc.Message {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	if len(n.listenChannels) == 0 {
		n.init()
	}
	c, ok := n.listenChannels[id]
	if !ok {
		return n.closedListenChan
	}
	return c
}

// Sending a message to the network.
func (n *OnlineNetwork) Send(curr *Node, msg *mpc.Message) {
	bz, err := msg.MarshalBinary()
	if err != nil {
		fmt.Printf("error while marshaling message: %e", err)
		return
	}

	n.mtx.Lock()
	defer n.mtx.Unlock()

	for _, pn := range n.nodes {
		if msg.IsFor(pn.partyId) {
			err = curr.Publish(topicKey(pn.partyId), bz)
			if err != nil {
				fmt.Printf("error while publishing message: %e", err)
				continue
			}
		}
	}
}

// Closing the subscriptions and returning the done channel.
func (n *OnlineNetwork) Done(id party.ID) chan struct{} {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	for id, sub := range n.subscriptions {
		sub.Close()
		delete(n.listenChannels, id)
	}
	if len(n.listenChannels) == 0 {
		close(n.done)
	}
	return n.done
}

//
// Private methods
//

// A goroutine that is listening for messages on the topic handler.
func (n *OnlineNetwork) handleSubscription(id party.ID, sub icore.PubSubSubscription) error {
	for {
		msg, err := sub.Next(context.Background())
		if err != nil {
			return err
		}
		m := &mpc.Message{}
		err = m.UnmarshalBinary(msg.Data())
		if err != nil {
			return err
		}
		n.listenChannels[id] <- m
	}
}

// topicKey returns the modified topic key for the given party ID.
func topicKey(id party.ID) string {
	return fmt.Sprintf("%s/mpc/cmp-keygen", string(id))
}
