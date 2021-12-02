package app

import (
	"flag"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/camera/internal/config"
	"github.com/defaulterrr/elegant_swirles/camera/internal/server"
	"github.com/defaulterrr/elegant_swirles/camera/internal/service"
)

func Run() error {
	var pathToConfig string

	flag.StringVar(&pathToConfig, "config", "./config.yaml", "Specify a path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(pathToConfig)
	if err != nil {
		return fmt.Errorf("config.NewConfig: %v", err)
	}

	newService := service.NewService()

	if err := server.NewServer(newService).Start(&cfg.Grpc); err != nil {
		return fmt.Errorf("server.Start: %v", err)
	}

	return nil
}
