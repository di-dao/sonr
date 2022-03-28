package channel

import (
	"context"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	v1 "go.buf.build/grpc/go/sonr-io/core/host/channel/v1"
	nh "github.com/sonr-io/core/host"
	"google.golang.org/protobuf/proto"
)

var (
	logger            *golog.Logger
	ErrNotOwner       = errors.New("Not owner of key - (Beam)")
	ErrNotFound       = errors.New("Key not found in store - (Beam)")
	ErrInvalidMessage = errors.New("Invalid message received in Pubsub Topic - (Beam)")
)

// Option is a function that modifies the beam options.
type Option func(*options)

// WithTTL sets the time-to-live for the beam store entries
func WithTTL(ttl time.Duration) Option {
	return func(o *options) {
		o.ttl = ttl
	}
}

// WithCapacity sets the capacity of the beam store.
func WithCapacity(capacity int) Option {
	return func(o *options) {
		o.capacity = capacity
	}
}

// options is a collection of options for the beam.
type options struct {
	ttl      time.Duration
	capacity int
}

// defaultOptions is the default options for the beam.
func defaultOptions() *options {
	return &options{
		ttl:      time.Minute * 10,
		capacity: 4096,
	}
}

// Channel is a pubsub based Key-Value store for Libp2p nodes.
type Channel interface {
	// Did returns the DID of the channel.
	Did() string

	// Get returns the value for the given key.
	Get(key string) ([]byte, error)

	// Put stores the value for the given key.
	Put(key string, value []byte) error

	// Delete removes the value for the given key.
	Delete(key string) error

	// Read returns a list of all peers subscribed to the channel topic.
	Read() []peer.ID

	// Publish publishes the given message to the channel topic.
	Publish(text string, data []byte) error

	// Listen subscribes to the beam topic and returns a channel that will
	// receive events.
	Listen() (<-chan *v1.ChannelMessage, error)

	// Close closes the channel.
	Close() error
}

// channel is the implementation of the Beam interface.
type channel struct {
	Channel
	ctx   context.Context
	n     nh.HostImpl
	label string
	did   string

	// Channel Messages
	messages        chan *v1.ChannelMessage
	messagesHandler *ps.TopicEventHandler
	messagesSub     *ps.Subscription
	messagesTopic   *ps.Topic

	// Store Events
	storeEvents        chan *v1.ChannelEvent
	storeEventsHandler *ps.TopicEventHandler
	storeEventsSub     *ps.Subscription
	storeEventsTopic   *ps.Topic
	store              *v1.ChannelStore
}

// New creates a new beam with the given name and options.
func New(ctx context.Context, n nh.HostImpl, id string, options ...Option) (Channel, error) {
	logger = golog.Default.Child(id)
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	mTopic, mHandler, mSub, err := n.NewTopic(id)
	if err != nil {
		return nil, err
	}

	evTopic, evHandler, evSub, err := n.NewTopic(id)
	if err != nil {
		return nil, err
	}

	b := &channel{
		ctx:                ctx,
		n:                  n,
		did:                id,
		messages:           make(chan *v1.ChannelMessage),
		messagesHandler:    mHandler,
		messagesSub:        mSub,
		messagesTopic:      mTopic,
		storeEvents:        make(chan *v1.ChannelEvent),
		storeEventsTopic:   evTopic,
		storeEventsSub:     evSub,
		storeEventsHandler: evHandler,
		store:              NewStore(opts),
	}

	// Start the event handler.
	go b.handleChannelEvents()
	go b.handleChannelMessages()
	go b.handleStoreEvents()
	go b.handleStoreMessages()
	go b.serve()
	return b, nil
}

// Read lists all peers subscribed to the beam topic.
func (b *channel) Read() []peer.ID {
	storePeers := b.storeEventsTopic.ListPeers()
	messagesPeers := b.messagesTopic.ListPeers()

	// filter out duplicates
	peers := make(map[peer.ID]struct{})
	for _, p := range storePeers {
		peers[p] = struct{}{}
	}
	for _, p := range messagesPeers {
		peers[p] = struct{}{}
	}

	// convert to slice
	var result []peer.ID
	for p := range peers {
		result = append(result, p)
	}
	return result
}

// Publish publishes the given message to the beam topic.
func (b *channel) Publish(text string, data []byte) error {
	// Check if both text and data are empty.
	if text == "" && len(data) == 0 {
		return errors.New("text and data cannot be empty")
	}

	// Create the message.
	msg := &v1.ChannelMessage{
		Text:  text,
		Data:  data,
		Owner: b.n.HostID().String(),
	}

	// Encode the message.
	buf, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	// Publish the message to the beam topic.
	return b.messagesTopic.Publish(b.ctx, buf)
}

// Delete removes the key in the beam store.
func (b *channel) Delete(key string) error {
	return DeleteStoreKey(b.store, key, b)
}

// Get returns the value for the given key in the beam store.
func (b *channel) Get(key string) ([]byte, error) {
	return GetKey(b.store, key)
}

// Put stores the value for the given key in the beam store.
func (b *channel) Put(key string, value []byte) error {
	return PutStoreKey(b.store, key, value, b)
}

// Listen subscribes to the beam topic and returns a channel that will
func (b *channel) Listen() (<-chan *v1.ChannelMessage, error) {
	return b.messages, nil
}

// Close closes the channel.
func (b *channel) Close() error {
	err := b.storeEventsTopic.Close()
	if err != nil {
		return err
	}

	err = b.messagesTopic.Close()
	if err != nil {
		return err
	}
	return nil
}
