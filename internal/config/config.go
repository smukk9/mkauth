package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
}

type Server struct {
	Port int
	Mode string
}

func Load(configFile string) (*Config, error) {
	v := viper.New()
	var c Config
	v.SetConfigType("yaml")
	v.AddConfigPath(configFile)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config err: %w", err)
	}

	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("Unmarshal err: %w", err)
	}

	return &c, nil
}
