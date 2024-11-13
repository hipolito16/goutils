package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AES struct {
	Key []byte
}

// NewAES cria uma nova instância de AES com a chave fornecida.
func NewAES(key string) (*AES, error) {
	if len(key) != 32 {
		return nil, errors.New("A chave deve ter 32 bytes para AES-256")
	}

	return &AES{Key: []byte(key)}, nil
}

// Encrypt criptografa uma string com AES-GCM.
func (self *AES) Encrypt(value string) (string, error) {
	block, err := aes.NewCipher(self.Key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	encrypted := aesgcm.Seal(iv, iv, []byte(value), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt descriptografa uma string com AES-GCM.
func (self *AES) Decrypt(value string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(self.Key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	iv, ciphertext := data[:aesgcm.NonceSize()], data[aesgcm.NonceSize():]
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
