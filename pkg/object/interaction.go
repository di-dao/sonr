package object

import (
	"fmt"

	"github.com/google/uuid"
)

func (ao *AppObjectInternalImpl) CreateObject(
	label string,
	object map[string]interface{}) (*ObjectUploadResult, error) {
	err := ao.schemaInternal.VerifyObject(object)

	if err != nil {
		return nil, err
	}

	err = ao.schemaInternal.BuildNodesFromDefinition(object)
	if err != nil {
		return nil, err
	}

	enc, err := ao.schemaInternal.EncodeDagJson()

	if err != nil {
		return nil, err
	}

	cid, err := ao.shell.DagPut(enc, "dag-json", "dag-cbor")
	did := fmt.Sprintf("did:snr:%s", uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &ObjectUploadResult{
		Code: 200,
		Definition: &ObjectDefinition{
			Did:   did,
			Cid:   cid,
			Label: label,
		},
		Message: "Object uploaded",
	}, nil
}

func (ao *AppObjectInternalImpl) GetObject(cid string) (map[string]interface{}, error) {
	var dag map[string]interface{}
	err := ao.shell.DagGet(cid, &dag)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return dag, err
}
