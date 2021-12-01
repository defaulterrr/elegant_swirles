package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/defaulterrr/elegant_swirles/processing/internal/config"
	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
	"github.com/defaulterrr/elegant_swirles/processing/internal/repository"
	"github.com/defaulterrr/elegant_swirles/processing/internal/server"
	"github.com/defaulterrr/elegant_swirles/processing/internal/service"
	"github.com/defaulterrr/elegant_swirles/processing/metrics"
)

func Run() {
	var pathToConfig string

	flag.StringVar(&pathToConfig, "config", "./config.yaml", "Specify a path to config file")
	flag.Parse()

	config, err := config.NewConfig(pathToConfig)
	if err != nil {
		fmt.Printf("config.NewConfig: %v", err)
		os.Exit(1)
	}

	err = metrics.InitMetrics()
	if err != nil {
		fmt.Printf("metrics.InitMetrics: %v", err)
		os.Exit(1)
	}

	conns, err := repository.GetGRPCConns(&config.Grpc)
	if err != nil {
		log.Fatalf("repository.GetGRPCConns: %v", err)
	}

	defer func() {
		err = conns.DHTConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	repo := repository.NewRepository(conns)
	serv := service.NewService(repo)

	// testing GetDHTMetrics function
	metricsServer := server.CreateMetricsServer(&config.Metrics)

	go func() {
		fmt.Printf("Metrics server is running on %s:%d\n", config.Metrics.Host, config.Metrics.Port)
		if err = metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Failed running metrics server:%v\n", err)
			os.Exit(1)
		}
	}()

	curMetrics := make(chan model.DHTMetrics)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err = serv.GetDHTMetrics(ctx, curMetrics)
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
}
