package kms

import (
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

type Service interface {
	/*CreateBinanceSignature(ctx context.Context, payload, apiSecret string) string
	CreateBinanceSignatureByAccountID(ctx context.Context, payload, accountID string) (string, string, error)*/
	Encrypt(plaintext string) (string, error)
	Hash(plaintext string) string
	Decrypt(encoded string) (string, error)
}

type service struct {
	gcm     cipher.AEAD
	hmacKey []byte
}

func newService(gcm cipher.AEAD, hmacKey []byte) Service {
	return service{gcm, hmacKey}
}

func (s service) Encrypt(plaintext string) (string, error) {
	nonce := make([]byte, s.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := s.gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (s service) Hash(plaintext string) string {
	mac := hmac.New(sha256.New, s.hmacKey)
	mac.Write([]byte(plaintext))
	return hex.EncodeToString(mac.Sum(nil))
}

func (s service) Decrypt(encoded string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	ns := s.gcm.NonceSize()
	if len(data) < ns {
		return "", errors.New("ciphertext too short")
	}
	nonce, ct := data[:ns], data[ns:]
	plaintext, err := s.gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
