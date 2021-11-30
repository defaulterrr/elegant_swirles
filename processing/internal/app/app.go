package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/defaulterrr/iot3/processing/internal/config"
	"github.com/defaulterrr/iot3/processing/internal/model"
	"github.com/defaulterrr/iot3/processing/internal/repository"
	"github.com/defaulterrr/iot3/processing/internal/service"
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

	conns, err := repository.GetGRPCConns(&config.Grpc)
	if err != nil {
		log.Fatalf("repository.GetGRPCConns: %v", err)
	}

	defer conns.DHTConn.Close()

	repo := repository.NewRepository(conns)
	serv := service.NewService(repo)

	// testing GetDHTMetrics function
	metrics := make(chan model.DHTMetrics)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err = serv.GetDHTMetrics(ctx, metrics)
		if err != nil {
			fmt.Printf("GetDHTMetrics: %v", err)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(<-metrics)
	}

	cancel()

	fmt.Println("finish")
}
