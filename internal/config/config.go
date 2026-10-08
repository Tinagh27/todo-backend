package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string
}

func Load() (*Config, error) {
	v := viper.New()

	v.SetDefault("APP_PORT", "8080")

	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		var pathErr *os.PathError
		if !errors.As(err, &pathErr) {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}

	return &Config{
		AppPort: v.GetString("APP_PORT"),
	}, nil
}
