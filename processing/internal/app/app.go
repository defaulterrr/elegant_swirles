package app

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/defaulterrr/elegant_swirles/processing/internal/config"
	"github.com/defaulterrr/elegant_swirles/processing/internal/repository"
	"github.com/defaulterrr/elegant_swirles/processing/internal/server"
	"github.com/defaulterrr/elegant_swirles/processing/internal/service"
	"github.com/defaulterrr/elegant_swirles/processing/metrics"
)

func Run() error {
	var pathToConfig string

	flag.StringVar(&pathToConfig, "config", "./config.yaml", "Specify a path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(pathToConfig)
	if err != nil {
		return fmt.Errorf("config.NewConfig: %v", err)
	}

	err = metrics.InitMetrics()
	if err != nil {
		return fmt.Errorf("metrics.InitMetrics: %v", err)
	}

	conns, err := repository.GetGRPCConns(cfg)
	if err != nil {
		return fmt.Errorf("repository.GetGRPCConns: %v", err)
	}

	defer func() {
		err = conns.DHTConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	newRepo := repository.NewRepository(conns)
	newService := service.NewService(newRepo)
	newServer := server.NewServer(newService)

	err = newServer.Start(&cfg.Metrics)
	if err != nil {
		fmt.Printf("Start: %v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := newServer.ReadDHTMetrics(ctx)
		if err != nil {
			fmt.Printf("ReadDHTMetrics: %v", err)
		}
	}()

	go func() {
		err := newServer.ReadCameraMetrics(ctx)
		if err != nil {
			fmt.Printf("ReadCameraMetrics: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	cancel()

	fmt.Println("finish")

	return nil
}
