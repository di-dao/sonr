package common

import (
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// NewLastUpdated returns a new LastUpdated Object as int64
var NewLastUpdated = func() int64 {
	return time.Now().Unix()
}

// RPC_SERVER_PORT is the port the RPC service listens on.
// Calculated: (Sister Bday + Dad Bday + Mom Bday) / Mine
const RPC_SERVER_PORT = 26225

// GetProfileFunc returns a function that returns the Profile and error
type GetProfileFunc func() (*Profile, error)

var (
	logger = golog.Child("internal/common")
)

// IsMdnsCompatible returns true if the Connection is MDNS compatible
func (c Connection) IsMdnsCompatible() bool {
	return c == Connection_WIFI || c == Connection_ETHERNET
}

// Checks if Enviornment is Development
func (e Environment) IsDev() bool {
	return e == Environment_DEVELOPMENT
}

// Checks if Enviornment is Development
func (e Environment) IsProd() bool {
	return e == Environment_PRODUCTION
}

// WrapErrors wraps errors list into a single error
func WrapErrors(msg string, errs []error) error {
	// Check if errors are empty
	if len(errs) == 0 {
		return nil
	}

	// Iterate over errors
	err := errors.New(msg)
	for _, e := range errs {
		if e != nil {
			err = errors.Wrap(e, e.Error())
			continue
		}
	}
	return err
}

func DefaultLocation() *Location {
	return &Location{
		Latitude:  34.102920,
		Longitude: -118.394190,
	}
}

func (l *Location) OLC() string {
	return olc.Encode(l.GetLatitude(), l.GetLongitude(), 6)
}

// ** ───────────────────────────────────────────────────────
// ** ─── Payload Management ────────────────────────────────
// ** ───────────────────────────────────────────────────────
// PayloadItemFunc is the Map function for PayloadItem
type PayloadItemFunc func(item *Payload_Item, index int, total int) error

// IsSingle returns true if the transfer is a single transfer. Error returned
// if No Items present in Payload
func (p *Payload) IsSingle() (bool, error) {
	if len(p.GetItems()) == 0 {
		return false, errors.New("No Items present in Payload")
	}
	if len(p.GetItems()) > 1 {
		return false, nil
	}
	return true, nil
}

// IsMultiple returns true if the transfer is a multiple transfer. Error returned
// if No Items present in Payload
func (p *Payload) IsMultiple() (bool, error) {
	if len(p.GetItems()) == 0 {
		return false, errors.New("No Items present in Payload")
	}
	if len(p.GetItems()) > 1 {
		return true, nil
	}
	return false, nil
}

// FileCount returns the number of files in the Payload
func (p *Payload) FileCount() int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type != MIME_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
}

// MapItems performs method chaining on the Items in the Payload
func (p *Payload) MapItems(fn PayloadItemFunc) error {
	count := len(p.GetItems())
	for i, item := range p.GetItems() {
		if err := fn(item, i, count); err != nil {
			return err
		}
	}
	return nil
}

// MapItems performs method chaining on the Items in the Payload
func (p *Payload) MapItemsWithIndex(fn PayloadItemFunc) error {
	count := len(p.GetItems())
	for i, item := range p.GetItems() {
		if err := fn(item, i, count); err != nil {
			return err
		}
	}
	return nil
}

// MapFileItems performs method chaining on ONLY the FileItems in the Payload
func (p *Payload) MapFileItems(fn PayloadItemFunc) error {
	count := len(p.GetItems())
	for i, item := range p.GetItems() {
		if item.GetFile() != nil {
			if err := fn(item, i, count); err != nil {
				return err
			}
		}
	}
	return nil
}

// MapUrlItems performs method chaining on ONLY the UrlItems in the Payload
func (p *Payload) MapUrlItems(fn PayloadItemFunc) error {
	count := len(p.GetItems())
	for i, item := range p.GetItems() {
		if item.GetUrl() != nil {
			if err := fn(item, i, count); err != nil {
				return err
			}
		}
	}
	return nil
}

// ReplaceItemsDir iterates over the items in the payload and replaces the
// directory of the item with the new directory.
func (p *Payload) ResetItemsDirectory() *Payload {
	// Create new Payload
	for _, item := range p.GetItems() {
		if item.GetFile() != nil {
			item.GetFile().ResetDir()
		}
	}
	return p
}

// URLCount returns the number of URLs in the Payload
func (p *Payload) URLCount() int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type == MIME_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
}

// Buffer returns Peer as a buffer
func (p *Profile) Buffer() ([]byte, error) {
	// Marshal Peer
	data, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}

	// Return Peer as buffer
	return data, nil
}
