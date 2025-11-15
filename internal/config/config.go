package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var ErrEnvVarNotSet = errors.New("config env MKAUTH_FILE is not set")

type Config struct {
	Server   Server
	Database Database
	Admin    Admin
}

type Server struct {
	Port    int
	Host    string
	Mode    string
	Version string
	Service string
}

type Database struct {
	Path string
}

type Admin struct {
	Email    string
	Password string
}

func Load() (*Config, error) {
	v := viper.New()
	configFile := os.Getenv("MKAUTH_FILE")

	if configFile == "" {
		return nil, ErrEnvVarNotSet
	}
	var c Config

	v.SetConfigFile(configFile)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("Unmarshal err: %w", err)
	}

	return &c, nil
}
