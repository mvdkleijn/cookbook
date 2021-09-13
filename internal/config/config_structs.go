package config

import "gorm.io/gorm"

type Config struct {
	Global         GlobalConfig
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
