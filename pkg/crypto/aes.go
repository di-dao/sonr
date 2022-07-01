package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

// AesEncryptWithKey uses the give 32-bit key to encrypt plaintext.
func AesEncryptWithKey(aesKey, plaintext []byte) (string, error) {
	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return string(ciphertext), nil
}

// AesDecryptWithKey uses the give 32-bit key to decrypt plaintext.
func AesDecryptWithKey(aesKey, ciphertext []byte) (string, error) {
	blockCipher, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return "", err
	}

	nonce, ct := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// AesDecryptWithPassword uses the give password to generate an aes key and decrypt plaintext.
func AesEncryptWithPassword(password string, plaintext []byte) (string, error) {
	key, err := deriveKey(password)
	if err != nil {
		return "", err
	}

	return AesEncryptWithKey(key, plaintext)
}

// AesEncryptWithPassword uses the give password to generate an aes key and encrypt plaintext.
func AesDecryptWithPassword(password string, ciphertext []byte) (string, error) {
	key, err := deriveKey(password)
	if err != nil {
		return "", err
	}

	return AesDecryptWithKey(key, ciphertext)
}

func deriveKey(password string) ([]byte, error) {
	// including a salt would make it impossible to reliably login from other devices
	key, err := scrypt.Key([]byte(password), []byte(""), 1<<20, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// NewAesKey generates a new 32-bit key.
func NewAesKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}
