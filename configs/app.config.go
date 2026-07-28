package configs

import (
	"p2ptrader/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AppConfig struct {
	Profile                 string
	Environment             string
	PocketbaseAdminEmail    string
	PocketbaseAdminPassword string
	GOMEMLIMIT              string
	RateLimitEnabled        string
	RateLimitRules          string
	PocketbaseEncryptionKey string
	//PocketbaseOauthTGBotToken string
	PocketbaseApiPrefix    string
	AesEncryptionMasterKey string
	TTlMemoryCacheDuration string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		Profile:                 utils.Getenv("PROFILE"),
		Environment:             utils.Getenv("ENVIRONMENT"),
		PocketbaseAdminEmail:    utils.Getenv("POCKETBASE_ADMIN_EMAIL", "admin@rublex.exchange"),
		PocketbaseAdminPassword: utils.Getenv("POCKETBASE_ADMIN_PASSWORD", "adminadmin"),
		PocketbaseEncryptionKey: utils.Getenv("POCKETBASE_ENCRYPTION_KEY"),
		PocketbaseApiPrefix:     utils.Getenv("POCKETBASE_API_PREFIX", "/api/v1"),
		GOMEMLIMIT:              utils.Getenv("GOMEMLIMIT"),
		RateLimitEnabled:        utils.Getenv("RATE_LIMIT_ENABLED"),
		RateLimitRules:          utils.Getenv("RATE_LIMIT_RULES"),
		AesEncryptionMasterKey:  utils.Getenv("AES_ENCRYPTION_MASTER_KEY"),
		TTlMemoryCacheDuration:  utils.Getenv("MAX_TTL_MEMORY_CACHE_DURATION", "18000"),
	}
}

func (c AppConfig) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.AesEncryptionMasterKey, validation.Required, validation.Length(64, 64)),
		validation.Field(&c.PocketbaseAdminEmail, is.EmailFormat),
		validation.Field(&c.PocketbaseAdminPassword, validation.Length(8, 0)),
		validation.Field(&c.PocketbaseEncryptionKey, validation.Length(32, 32)),
	)
}
