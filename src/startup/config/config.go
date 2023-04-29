package config

import "os"

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBName string
}

func NewLocalConfig() *Config {
	return &Config{
		Port:   os.Getenv("SERVER_PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
}
