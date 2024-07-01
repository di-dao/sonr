package did

import (
	"fmt"

	"github.com/di-dao/sonr/internal/models"
)

// Format.Credential formats a credential as a DID
func (f formatter) Credential(c *models.Credential) string {
	return fmt.Sprintf("did:web:%s@%s#%s", c.DisplayName, c.Origin, c.ID)
}

// Format.Wallet formats a wallet as a DID
func (f formatter) Wallet(w *models.Wallet) string {
	return fmt.Sprintf("did:sonr:%s", w.Address)
}

// Format is the formatter to use for DIDs
var Format = formatter{}

type formatter struct{}
