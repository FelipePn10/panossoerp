package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
}

func Load() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Defaults (fallback)
	viper.SetDefault("SERVER_ADDR", "5070")
	viper.SetDefault(
		"DATABASE_URL",
		"postgres://panossoerp:panossoerp_10203040@localhost:5432/panossoerpdatabase?sslmode=disable",
	)

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	return &Config{
		ServerPort:  viper.GetString("SERVER_ADDR"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
	}
}
