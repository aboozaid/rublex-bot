package pocketbase

import "os"

type Config struct {
	Profile                   string
	Environment               string
	PocketbaseAdminEmail      string
	PocketbaseAdminPassword   string
	GOMEMLIMIT                string
	RateLimitEnabled          string
	RateLimitRules            string
	PocketbaseEncryptionKey   string
	PocketbaseOauthTGBotToken string
}

func GetConfig() Config {
	return Config{
		Profile:                   os.Getenv("PROFILE"),
		Environment:               os.Getenv("ENVIRONMENT"),
		PocketbaseAdminEmail:      os.Getenv("POCKETBASE_ADMIN_EMAIL"),
		PocketbaseAdminPassword:   os.Getenv("POCKETBASE_ADMIN_PASSWORD"),
		GOMEMLIMIT:                os.Getenv("GOMEMLIMIT"),
		RateLimitEnabled:          os.Getenv("RATE_LIMIT_ENABLED"),
		RateLimitRules:            os.Getenv("RATE_LIMIT_RULES"),
		PocketbaseEncryptionKey:   os.Getenv("POCKETBASE_ENCRYPTION_KEY"),
		PocketbaseOauthTGBotToken: os.Getenv("POCKETBASE_OAUTH_TG_BOT_TOKEN"),
	}
}
