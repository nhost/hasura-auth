package crypto_test

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"testing"

	"github.com/nhost/hasura-auth/go/lib/crypto"
)

func TestNewEncrypter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		keySize   int
		shouldErr bool
	}{
		{
			name:      "AES-128 key (16 bytes)",
			keySize:   16,
			shouldErr: false,
		},
		{
			name:      "AES-192 key (24 bytes)",
			keySize:   24,
			shouldErr: false,
		},
		{
			name:      "AES-256 key (32 bytes)",
			keySize:   32,
			shouldErr: false,
		},
		{
			name:      "invalid key size (8 bytes)",
			keySize:   8,
			shouldErr: true,
		},
		{
			name:      "invalid key size (0 bytes)",
			keySize:   0,
			shouldErr: true,
		},
		{
			name:      "invalid key size (33 bytes)",
			keySize:   33,
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			key := make([]byte, tt.keySize)
			if tt.keySize > 0 {
				_, err := rand.Read(key)
				if err != nil {
					t.Fatalf("failed to generate test key: %v", err)
				}
			}

			enc, err := crypto.NewEncrypter(key)

			if tt.shouldErr { //nolint:nestif
				if err == nil {
					t.Error("expected error but got none")
				}

				if enc != nil {
					t.Error("expected nil encrypter when error occurs")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if enc == nil {
					t.Error("expected valid encrypter but got nil")
				}
			}
		})
	}
}

func TestEncryptDecrypt(t *testing.T) {
	t.Parallel()

	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("failed to generate test key: %v", err)
	}

	enc, err := crypto.NewEncrypter(key)
	if err != nil {
		t.Fatalf("failed to create encrypter: %v", err)
	}

	tests := []struct {
		name      string
		plaintext []byte
	}{
		{
			name:      "empty data",
			plaintext: []byte{},
		},
		{
			name:      "single byte",
			plaintext: []byte{0x42},
		},
		{
			name:      "short text",
			plaintext: []byte("hello"),
		},
		{
			name: "long text",
			plaintext: []byte(
				"This is a longer piece of text that should be encrypted and decrypted successfully using AES-GCM encryption.",
			),
		},
		{
			name:      "binary data",
			plaintext: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0xFF, 0xFE, 0xFD},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ciphertext, err := enc.Encrypt(tt.plaintext)
			if err != nil {
				t.Errorf("encryption failed: %v", err)
				return
			}

			if len(ciphertext) == 0 {
				t.Error("ciphertext is empty")
				return
			}

			decrypted, err := enc.Decrypt(ciphertext)
			if err != nil {
				t.Errorf("decryption failed: %v", err)
				return
			}

			if !bytes.Equal(tt.plaintext, decrypted) {
				t.Errorf(
					"decrypted data doesn't match original\noriginal:  %v\ndecrypted: %v",
					tt.plaintext,
					decrypted,
				)
			}
		})
	}
}

func TestEncryptWithAAD(t *testing.T) { //nolint:cyclop
	t.Parallel()

	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("failed to generate test key: %v", err)
	}

	enc, err := crypto.NewEncrypter(key)
	if err != nil {
		t.Fatalf("failed to create encrypter: %v", err)
	}

	tests := []struct {
		name          string
		plaintext     []byte
		aad           []byte
		wrongAAD      []byte
		shouldDecrypt bool
	}{
		{
			name:          "with AAD",
			plaintext:     []byte("secret message"),
			aad:           []byte("user:123"),
			wrongAAD:      []byte("user:456"),
			shouldDecrypt: true,
		},
		{
			name:          "with empty AAD",
			plaintext:     []byte("another secret"),
			aad:           []byte{},
			wrongAAD:      []byte("not empty"),
			shouldDecrypt: true,
		},
		{
			name:          "with nil AAD",
			plaintext:     []byte("third secret"),
			aad:           nil,
			wrongAAD:      []byte("something"),
			shouldDecrypt: true,
		},
		{
			name:          "with complex AAD",
			plaintext:     []byte("complex secret"),
			aad:           []byte("session:abc123,user:456,role:admin"),
			wrongAAD:      []byte("session:abc123,user:456,role:user"),
			shouldDecrypt: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ciphertext, err := enc.EncryptWithAAD(tt.plaintext, tt.aad)
			if err != nil {
				t.Errorf("encryption with AAD failed: %v", err)
				return
			}

			if tt.shouldDecrypt {
				decrypted, err := enc.DecryptWithAAD(ciphertext, tt.aad)
				if err != nil {
					t.Errorf("decryption with correct AAD failed: %v", err)
					return
				}

				if !bytes.Equal(tt.plaintext, decrypted) {
					t.Errorf(
						"decrypted data doesn't match original\noriginal:  %v\ndecrypted: %v",
						tt.plaintext,
						decrypted,
					)
				}
			}

			if tt.wrongAAD != nil {
				_, err := enc.DecryptWithAAD(ciphertext, tt.wrongAAD)
				if err == nil {
					t.Error("expected decryption with wrong AAD to fail")
				}

				if !errors.Is(err, crypto.ErrDecryptionFailed) {
					t.Errorf("expected ErrDecryptionFailed, got: %v", err)
				}
			}
		})
	}
}

func TestDecryptErrors(t *testing.T) {
	t.Parallel()

	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("failed to generate test key: %v", err)
	}

	enc, err := crypto.NewEncrypter(key)
	if err != nil {
		t.Fatalf("failed to create encrypter: %v", err)
	}

	tests := []struct {
		name        string
		ciphertext  []byte
		expectedErr error
	}{
		{
			name:        "too short ciphertext",
			ciphertext:  []byte("dG9vIHNob3J0"), // base64 encoded "too short" (9 bytes)
			expectedErr: crypto.ErrCiphertextTooShort,
		},
		{
			name:        "empty ciphertext",
			ciphertext:  []byte{},
			expectedErr: nil, // Will fail at decode stage
		},
		{
			name:        "invalid base64",
			ciphertext:  []byte("not valid base64!@#$"),
			expectedErr: nil, // Will fail at decode stage with base64 error
		},
		{
			name: "valid length but corrupted data",
			ciphertext: func() []byte {
				// Create a base64-encoded string with enough length but corrupted auth tag
				// This simulates data that passes base64 decode but fails GCM authentication
				validCiphertext, _ := enc.Encrypt([]byte("test"))
				// Decode, corrupt last byte, re-encode
				decoded, _ := crypto.Decode(validCiphertext)
				if len(decoded) > 0 {
					decoded[len(decoded)-1] ^= 0x01
				}
				corrupted, _ := crypto.Encode(decoded)
				return corrupted
			}(),
			expectedErr: crypto.ErrDecryptionFailed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := enc.Decrypt(tt.ciphertext)
			if err == nil {
				t.Error("expected decryption to fail")
				return
			}

			if tt.expectedErr != nil && !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got: %v", tt.expectedErr, err)
			}
		})
	}
}

func TestGenerateKey(t *testing.T) { //nolint:cyclop
	t.Parallel()

	// Test multiple key generations
	keys := make([][]byte, 10)

	for i := range 10 {
		key, err := crypto.GenerateKey()
		if err != nil {
			t.Errorf("GenerateKey() failed: %v", err)
			continue
		}

		if len(key) != 32 {
			t.Errorf("expected key length 32, got: %d", len(key))
			continue
		}

		keys[i] = key
	}

	// Verify keys are different (probability of collision is astronomically low)
	for i := range keys {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] != nil && keys[j] != nil && bytes.Equal(keys[i], keys[j]) {
				t.Error("generated keys should be different")
			}
		}
	}

	// Test that generated key works with encrypter
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey() failed: %v", err)
	}

	enc, err := crypto.NewEncrypter(key)
	if err != nil {
		t.Errorf("NewEncrypter with generated key failed: %v", err)
	}

	if enc == nil {
		t.Error("encrypter should not be nil")
		return
	}

	// Test encryption/decryption with generated key
	plaintext := []byte("test message with generated key")

	ciphertext, err := enc.Encrypt(plaintext)
	if err != nil {
		t.Errorf("encryption with generated key failed: %v", err)
		return
	}

	decrypted, err := enc.Decrypt(ciphertext)
	if err != nil {
		t.Errorf("decryption with generated key failed: %v", err)
		return
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Error("round-trip with generated key failed")
	}
}

func TestEncryptionNonceUniqueness(t *testing.T) { //nolint:cyclop
	t.Parallel()

	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("failed to generate test key: %v", err)
	}

	enc, err := crypto.NewEncrypter(key)
	if err != nil {
		t.Fatalf("failed to create encrypter: %v", err)
	}

	plaintext := []byte("same message")
	ciphertexts := make([][]byte, 10)

	// Encrypt same message multiple times
	for i := range 10 {
		ciphertext, err := enc.Encrypt(plaintext)
		if err != nil {
			t.Errorf("encryption %d failed: %v", i, err)
			continue
		}

		ciphertexts[i] = ciphertext
	}

	// Verify all ciphertexts are different (due to random nonces)
	for i := range ciphertexts {
		for j := i + 1; j < len(ciphertexts); j++ {
			if ciphertexts[i] != nil && ciphertexts[j] != nil &&
				bytes.Equal(ciphertexts[i], ciphertexts[j]) {
				t.Error(
					"encrypting same plaintext should produce different ciphertexts due to random nonces",
				)
			}
		}
	}

	// Verify all can be decrypted to same plaintext
	for i, ciphertext := range ciphertexts {
		if ciphertext == nil {
			continue
		}

		decrypted, err := enc.Decrypt(ciphertext)
		if err != nil {
			t.Errorf("decryption %d failed: %v", i, err)
			continue
		}

		if !bytes.Equal(plaintext, decrypted) {
			t.Errorf("decryption %d produced wrong result", i)
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	t.Parallel()

	testData := [][]byte{
		{},
		{0x00},
		{0xFF},
		{0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
		[]byte("hello world"),
	}

	for i, data := range testData {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()

			encoded, err := crypto.Encode(data)
			if err != nil {
				t.Errorf("encode failed: %v", err)
				return
			}

			decoded, err := crypto.Decode(encoded)
			if err != nil {
				t.Errorf("decode failed: %v", err)
				return
			}

			if !bytes.Equal(data, decoded) {
				t.Errorf(
					"encode/decode round-trip failed\noriginal: %v\ndecoded:  %v",
					data,
					decoded,
				)
			}
		})
	}
}
