package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))

	return hex.EncodeToString(hasher.Sum(nil))
}

// encrypts the provided data using the provided encryptionKey
func Encrypt(data []byte, encryptionKey string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(CreateHash(encryptionKey)))
	if err != nil {
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)

	return cipherText, nil
}

// decrypts the provided data using the provided encryptionKey
func Decrypt(data []byte, encryptionKey string) ([]byte, error) {
	key := []byte(CreateHash(encryptionKey))

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
