package app

import (
	"context"
	"flag"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/internal/config"
	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
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

	conns, err := repository.GetGRPCConns(&cfg.Grpc)
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

	// testing GetDHTMetrics function
	err = newServer.Start(&cfg.Metrics)
	if err != nil {
		fmt.Printf("Start: %v\n", err)
	}

	curMetrics := make(chan model.DHTMetrics)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err = newService.GetDHTMetrics(ctx, curMetrics)
		if err != nil {
			fmt.Printf("GetDHTMetrics: %v\n", err)
		}
	}()

	for i := 0; i < 200; i++ {
		metr, ok := <-curMetrics
		if !ok {
			break
		}

		metrics.SetTemperature(float64(metr.Temperature))
		fmt.Println(metr)
	}

	cancel()

	fmt.Println("finish")

	return nil
}
