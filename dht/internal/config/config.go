package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Grpc Grpc `yaml:"grpc"`
}

type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

func NewConfig(filePath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return nil, fmt.Errorf("cleanenv.ReadConfig: %v", err)
	}

	return &cfg, nil
}
