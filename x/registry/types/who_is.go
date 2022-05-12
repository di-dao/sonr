package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/sonr-io/sonr/pkg/did"
)

// ContainsAlias checks if the alias is in the list of aliases of the whois
func (w *WhoIs) ContainsAlias(target string) bool {
	for _, a := range w.Alias {
		if a == target {
			return true
		}
	}
	return false
}

// ContainsController checks if the controller is in the list of controllers of the whois
func (w *WhoIs) ContainsController(target string) bool {
	// Validates DID String
	if _, err := did.ParseDID(target); err != nil {
		return false
	}

	// Checks if the controller is in the list of controllers
	for _, c := range w.Controllers {
		if c == target {
			return true
		}
	}
	return false
}

// UnmarshalDidDocument unmarshals the whois document into a DID document
func (w *WhoIs) UnmarshalDidDocument() (*did.Document, error) {
	doc := did.Document{}
	err := doc.UnmarshalJSON(w.DidDocument)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

// CopyFromDidDocument copies the DID document into the WhoIs object if the DID document has the same DID owner
func (w *WhoIs) CopyFromDidDocument(doc *did.Document) error {
	if w.Owner != strings.TrimLeft(doc.ID.ID, "did:snr:") {
		return fmt.Errorf("owner mismatch: %s != %s", w.Owner, doc.ID.ID)
	}
	if doc.AlsoKnownAs != nil {
		w.Alias = doc.AlsoKnownAs
	}
	w.Controllers = doc.ControllersAsString()
	w.Timestamp = time.Now().Unix()
	w.IsActive = true
	return nil
}
