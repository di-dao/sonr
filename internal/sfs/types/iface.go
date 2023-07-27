package types

import "github.com/sonrhq/core/internal/crypto"

type SFSMap interface {
	// Add adds an item to the map
	Add(key string, value string) error

	// Remove removes an item from the map
	Remove(key string) error

	// Contains checks if an item is in the map
	Contains(key string) (bool, error)

	// GetAll returns all items in the map
	GetAll() (map[string]string, error)

	// Get returns the item at the given key
	Get(key string) (string, error)
}

type SFSSet interface {
	// Add adds an item to the set
	Add(item string) error

	// Remove removes an item from the set
	Remove(item string) error

	// Contains checks if an item is in the set
	Contains(item string) (bool, error)

	// GetAll returns all items in the set
	GetAll() ([]string, error)

	// Get returns the item at the given index
	Get(index int) (string, error)
}

type ZKSet interface {
	// Add adds an item to the set
	Add(item string) error

	// Remove removes an item from the set
	Remove(item string) error

	// Contains checks if an item is in the set
	Contains(item string) (bool, error)

	// Decrypt decrypts the set value
	Decrypt(key *crypto.Secp256k1PubKey, ciphertext []byte) ([]byte, error)

	// Encrypt encrypts the set value
	Encrypt(key *crypto.Secp256k1PubKey, bz []byte) ([]byte, error)
}
