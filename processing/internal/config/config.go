package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Grpc Grpc `yaml:"grpc"`
}

type Grpc struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func NewConfig(filePath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return nil, fmt.Errorf("cleanenv.ReadConfig: %v", err)
	}

	return &cfg, nil
}
