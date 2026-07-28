package kms

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log/slog"
	"os"
	"p2ptrader/common"
	"p2ptrader/modules/config"
	"sync"
)

type module struct {
	service Service
}

var (
	instance *module
	once     sync.Once
)

func NewModule() common.Module[Service] {
	once.Do(func() {
		configService := config.NewModule().GetService()
		masterKey := configService.App().AesEncryptionMasterKey
		if len(masterKey) != 64 {
			slog.Error("unable to parse AES_ENCRYPTION_MASTER_KEY variable, must be 64 hexadecimal string")
			os.Exit(1)
		}
		key, err := hex.DecodeString(masterKey)
		if err != nil {
			slog.Error("unable to parse master key", "error", err)
			os.Exit(1)
		}
		block, err := aes.NewCipher(key)
		if err != nil {
			slog.Error("unable to create cipher block", "error", err)
			os.Exit(1)
		}
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			slog.Error("unable to create gcm", "error", err)
			os.Exit(1)
		}
		instance = &module{newService(gcm, key)}
	})

	return instance
}

func (m *module) GetService() Service {
	return m.service
}
