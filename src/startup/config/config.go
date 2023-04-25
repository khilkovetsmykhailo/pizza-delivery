package config

import "os"

type Config struct {
	Port string
}

func NewLocalConfig() *Config {
	return &Config{
		Port: os.Getenv("SERVER_PORT"),
	}
}
