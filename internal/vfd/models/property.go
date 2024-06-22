package models

import (
	"errors"
	"fmt"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/accumulator"
	"github.com/di-dao/sonr/crypto/secret"
)

type Properties map[string]*accumulator.Accumulator

func NewProperties() Properties {
	return make(Properties)
}

type Property struct {
	Key        string `json:"key"`
	Controller string `json:"controller"`
	Value      []byte `json:"value"`
}

// Check validates the witness
func (p Properties) Check(publicKey crypto.PublicKey, key string, witness []byte) bool {
	sk, err := secret.NewKey(key, publicKey)
	if err != nil {
		return false
	}
	acc, ok := p[key]
	if !ok {
		return false
	}
	wit := &accumulator.MembershipWitness{}
	err = wit.UnmarshalBinary(witness)
	if err != nil {
		return false
	}
	return sk.VerifyWitness(acc, wit) == nil
}

// Set sets the property for the controller
func (p Properties) Set(publicKey crypto.PublicKey, key, value string) ([]byte, error) {
	sk, err := secret.NewKey(key, publicKey)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to get secret key"))
	}
	acc, err := sk.CreateAccumulator(value)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create accumulator"))
	}
	p[key] = acc
	witness, err := sk.CreateWitness(acc, value)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to create witness"))
	}
	return witness.MarshalBinary()
}

// Remove unlinks the property from the controller
func (p Properties) Remove(publicKey crypto.PublicKey, key, value string) error {
	sk, err := secret.NewKey(key, publicKey)
	if err != nil {
		return err
	}
	acc, ok := p[key]
	if !ok {
		return fmt.Errorf("property not found")
	}
	witness, err := sk.CreateWitness(acc, value)
	if err != nil {
		return err
	}
	// no need to continue if the property is not linked
	err = sk.VerifyWitness(acc, witness)
	if err != nil {
		return nil
	}
	newAcc, err := sk.UpdateAccumulator(acc, []string{}, []string{value})
	if err != nil {
		return err
	}
	p[key] = newAcc
	return nil
}

// Marshal the properties to a byte map
func (p Properties) Marshal() (map[string][]byte, error) {
	results := make(map[string][]byte, len(p))
	for k, v := range p {
		b, err := v.MarshalBinary()
		if err != nil {
			return nil, err
		}
		results[k] = b
	}
	return results, nil
}

// Unmarshal the properties from a byte map
func (p Properties) Unmarshal(m map[string][]byte) error {
	for k, v := range m {
		acc := &accumulator.Accumulator{}
		err := acc.UnmarshalBinary(v)
		if err != nil {
			return err
		}
		p[k] = acc
	}
	return nil
}
