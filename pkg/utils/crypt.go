package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"gopos/internal/config"
)

var (
	key  []byte
	once sync.Once
)

func loadKey() []byte {
	once.Do(func() {
		config.LoadEnv()

		k := os.Getenv("ENCRYPT_KEY")

		if len(k) != 32 {
			log.Fatalf("ENCRYPT_KEY must be exactly 32 bytes")
		}
		key = []byte(k)
	})
	return key
}

func Encrypt(plaintext string) (string, error) {
	key := loadKey()

	// AES-GCM setup
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the plaintext
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	encryptedText := base64.StdEncoding.EncodeToString(ciphertext)

	// Encode key as base64
	base64Key := base64.StdEncoding.EncodeToString(key)

	// Combine base64(key) + "." + encryptedText
	combined := fmt.Sprintf("%s.%s", base64Key, encryptedText)

	// Encode whole string as base64
	final := base64.StdEncoding.EncodeToString([]byte(combined))

	return final, nil
}

func Decrypt(encoded string) (string, error) {
	// Step 1: Decode outer base64
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 token: %w", err)
	}

	decoded := string(decodedBytes)

	// Step 2: Split base64Key and encryptedText
	parts := strings.SplitN(decoded, ".", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid token format, expected base64Key.encryptedText")
	}

	base64Key := parts[0]
	encryptedTextBase64 := parts[1]

	// Step 3: Decode base64Key to get AES key
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64Key: %w", err)
	}
	if len(key) != 32 {
		return "", fmt.Errorf("invalid key length: expected 32 bytes")
	}

	// Step 4: Decode encryptedText
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedTextBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode encrypted text: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher block: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedBytes) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce := encryptedBytes[:nonceSize]
	ciphertext := encryptedBytes[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	return string(plaintext), nil
}
