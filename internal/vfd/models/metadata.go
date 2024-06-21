package models

import (
	"encoding/json"
)

// Metadata represents metadata of a resource
type Metadata struct {
	Creds      Credentials `json:"credentials"`
	Properties Properties  `json:"properties"`
	Address    string      `json:"address"`
}

func (i *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Metadata) Unmarshal(data []byte) error {
	return json.Unmarshal(data, i)
}
