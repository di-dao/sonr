package models

import (
	"encoding/json"

	"github.com/di-dao/sonr/crypto"
)

// Metadata represents metadata of a resource
type Metadata struct {
	Creds     Credentials      `json:"credentials"`
	Address   string           `json:"address"`
	PublicKey crypto.PublicKey `json:"publicKey"`
}

func (i *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Metadata) Unmarshal(data []byte) error {
	return json.Unmarshal(data, i)
}
