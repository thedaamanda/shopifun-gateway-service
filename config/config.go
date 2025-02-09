package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string
	BaseURLPath string
}

func LoadConfig() (*Config, error) {
	viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	config := &Config{
		AppPort:     viper.GetString("APP_PORT"),
		BaseURLPath: viper.GetString("BASE_URL_PATH"),
	}

	return config, nil
}

func WriteTimeout() time.Duration {
	return 10 * time.Second
}

func ReadTimeout() time.Duration {
	return 10 * time.Second
}
