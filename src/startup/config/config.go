package config

type Config struct {
	Host string
	Port string
}

func NewLocalConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: "8080",
	}
}
