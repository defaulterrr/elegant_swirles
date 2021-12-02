package app

import (
	"context"
	"flag"
	"fmt"
	"sync"

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

	// testing GetDHTMetrics function
	err = newServer.Start(&cfg.Metrics)
	if err != nil {
		fmt.Printf("Start: %v\n", err)
	}

	curMetrics := make(chan model.DHTMetrics)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err = newService.DHTService.GetDHTMetrics(ctx, curMetrics)
		if err != nil {
			fmt.Printf("GetDHTMetrics: %v\n", err)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			metr, ok := <-curMetrics
			if !ok {
				break
			}

			metrics.SetTemperature(float64(metr.Temperature))
			metrics.SetHumidity(float64(metr.Humidity))
			fmt.Println(metr)
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		cancel()
	}()

	curCameraMetrics := make(chan model.CameraMetrics)
	cameraCtx, cameraCancel := context.WithCancel(context.Background())
	defer cameraCancel()

	go func() {
		err = newService.CameraService.GetCameraMetrics(cameraCtx, curCameraMetrics)
		if err != nil {
			fmt.Printf("GetCameraMetrics: %v\n", err)
		}
	}()

	for i := 0; i < 1000; i++ {
		cameraMetr, ok := <-curCameraMetrics
		if !ok {
			break
		}

		metrics.SetCountPeople(float64(cameraMetr.CountPeople))
		fmt.Println(cameraMetr)
	}
	cameraCancel()

	fmt.Println("finish")

	return nil
}
