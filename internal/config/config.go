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
		if _, ok := errors.AsType[*os.PathError](err); !ok {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}

	return &Config{
		AppPort: v.GetString("APP_PORT"),
	}, nil
}
