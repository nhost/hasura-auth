package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

var (
	ErrCiphertextTooShort = errors.New("ciphertext too short")
	ErrDecryptionFailed   = errors.New("decryption failed")
)

type Encrypter struct {
	aead cipher.AEAD
}

// NewEncrypter creates a new encrypter using AES-GCM (AEAD mode)
//
// The key must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256.
//
// Important: Keys must be kept secret and should be generated using
// cryptographically secure random sources (see GenerateKey()) or derived
// from passwords using a key derivation function like Argon2, scrypt, or PBKDF2.
// Never use a password directly as a key.
func NewEncrypter(key []byte) (*Encrypter, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	return &Encrypter{
		aead: aead,
	}, nil
}

// Encrypt encrypts the plaintext using AES-GCM
// The result includes the nonce (first 12 bytes) followed by the ciphertext and authentication tag.
// The output is encoded in base64 for safe transport/storage.
func (enc *Encrypter) Encrypt(plaintext []byte) ([]byte, error) {
	return enc.EncryptWithAAD(plaintext, nil)
}

// EncryptWithAAD encrypts the plaintext using AES-GCM with Additional Authenticated Data
//
// The AAD is authenticated but not encrypted, useful for binding ciphertext to context
// (e.g., user ID, session ID, or other metadata that should be verified but not hidden).
// The AAD must be provided again during decryption for authentication to succeed.
//
// The result includes the nonce (first 12 bytes) followed by the ciphertext and authentication tag.
// The output is encoded in base64 for safe transport/storage.
func (enc *Encrypter) EncryptWithAAD(plaintext, additionalData []byte) ([]byte, error) {
	// Generate a random nonce
	nonce := make([]byte, enc.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt and authenticate in one step
	ciphertext := enc.aead.Seal(nonce, nonce, plaintext, additionalData)

	// Encode based on format
	return Encode(ciphertext)
}

// Decrypt decrypts the ciphertext using AES-GCM
// Expects ciphertext to be base64 encoded.
// Automatically verifies authentication tag.
func (enc *Encrypter) Decrypt(ciphertext []byte) ([]byte, error) {
	return enc.DecryptWithAAD(ciphertext, nil)
}

// DecryptWithAAD decrypts the ciphertext using AES-GCM with Additional Authenticated Data
//
// Expects ciphertext to be base64 encoded.
// The AAD must match exactly what was provided during encryption.
// If the AAD doesn't match, decryption will fail with ErrDecryptionFailed.
func (enc *Encrypter) DecryptWithAAD(ciphertext, additionalData []byte) ([]byte, error) {
	// Decode
	decoded, err := Decode(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	// Check minimum length (nonce + at least auth tag)
	if len(decoded) < enc.aead.NonceSize() {
		return nil, ErrCiphertextTooShort
	}

	// Extract nonce and ciphertext
	nonce := decoded[:enc.aead.NonceSize()]
	encryptedData := decoded[enc.aead.NonceSize():]

	// Decrypt and verify authentication tag in one step
	plaintext, err := enc.aead.Open(nil, nonce, encryptedData, additionalData)
	if err != nil {
		return nil, ErrDecryptionFailed
	}

	return plaintext, nil
}

// Encode encodes data in base64 format. Exposed for convenience and testing.
func Encode(data []byte) ([]byte, error) {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, data)

	return encoded, nil
}

// Decode decodes base64 encoded data. Exposed for convenience and testing.
func Decode(data []byte) ([]byte, error) {
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(data)))

	n, err := base64.StdEncoding.Decode(decoded, data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}

	return decoded[:n], nil
}

// GenerateKey generates a random 256-bit key for AES-256
//
// This generates a cryptographically secure random key suitable for use with AES-256.
// Keys should be stored securely and never hardcoded in source code.
func GenerateKey() ([]byte, error) {
	key := make([]byte, 32) //nolint:mnd

	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}

	return key, nil
}
