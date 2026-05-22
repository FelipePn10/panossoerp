package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string `mapstructure:"SERVER_ADDR"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	Env         string `mapstructure:"ENV"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`

	// SMTP — e-mail alerts (optional; leave blank to disable)
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     string `mapstructure:"SMTP_PORT"`
	SMTPUser     string `mapstructure:"SMTP_USER"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
	SMTPFrom     string `mapstructure:"SMTP_FROM"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Defaults (fallback)
	viper.SetDefault("SERVER_ADDR", "5070")
	viper.SetDefault(
		"DATABASE_URL",
		"postgres://panossoerp:panossoerp_10203040@localhost:5432/panossoerpdatabase?sslmode=disable",
	)
	viper.SetDefault("ENV", "development")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("SMTP_HOST", "")
	viper.SetDefault("SMTP_PORT", "587")
	viper.SetDefault("SMTP_USER", "")
	viper.SetDefault("SMTP_PASSWORD", "")
	viper.SetDefault("SMTP_FROM", "")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error: %w", err)
		}
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("erro parse config: %w", err)
	}
	return &cfg, nil
}
