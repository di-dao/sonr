package models

import (
	"encoding/json"
	"time"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/pkg/fs"
)

// Metadata represents metadata of a resource
type Metadata struct {
	Address    string           `json:"address"`
	PublicKey  crypto.PublicKey `json:"publicKey"`
	Name       string           `json:"name"`
	Size       int64            `json:"size"`
	Type       string           `json:"type"`
	Created    time.Time        `json:"created"`
	Modified   time.Time        `json:"modified"`
	Owner      string           `json:"owner"`
	Permission string           `json:"permission"`
}

func (i *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Metadata) Unmarshal(data []byte) error {
	return json.Unmarshal(data, i)
}

func (i *Metadata) Save(file fs.File) error {
	bz, err := i.Marshal()
	if err != nil {
		return err
	}
	return file.Write(bz)
}
