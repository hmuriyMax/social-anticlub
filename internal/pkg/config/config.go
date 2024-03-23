package config

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const configKey = "config"

// Config struct for webapp config
type Config struct {
	PG struct {
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		DB       string `yaml:"db"`
		User     string `yaml:"user"`
		Pass     string `yaml:"pass"`
		PoolSize int    `yaml:"pool"`
	} `yaml:"pg"`

	Server struct {
		HTTPPort string `yaml:"http_port"`
		GRPCPort string `yaml:"grpc_port"`
	} `yaml:"server"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	d := yaml.NewDecoder(file)

	if err = d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func SetToCtx(ctx context.Context, config *Config) context.Context {
	return context.WithValue(ctx, configKey, config)
}

func GetFromCtx(ctx context.Context) *Config {
	return ctx.Value(configKey).(*Config)
}
