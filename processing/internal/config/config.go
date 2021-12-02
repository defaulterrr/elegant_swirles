package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DHTGrpc    DHTGrpc    `yaml:"dht_grpc"`
	CameraGrpc CameraGrpc `yaml:"camera_grpc"`
	Metrics    Metrics    `yaml:"metrics"`
}

type DHTGrpc struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type CameraGrpc struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Metrics struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

func NewConfig(filePath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return nil, fmt.Errorf("cleanenv.ReadConfig: %v", err)
	}

	return &cfg, nil
}
