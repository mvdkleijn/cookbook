package config

import "gorm.io/gorm"

type Config struct {
	Global         GlobalConfig
	Oidc           OidcConfig
	Database       DatabaseConfig
	DatabaseClient *gorm.DB
}

// GlobalConfig holds global configuration items
type GlobalConfig struct {
	Debug bool
}

// DatabaseConfig holds database configuration items
type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Database string
	Port     int
	SSLMode  string
	Timezone string
}

type OidcConfig struct {
	URL           string
	ClientID      string
	SigningAlgs   []string
	ClientIDCheck bool
	ExpiryCheck   bool
	IssuerCheck   bool
}
