package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/defaulterrr/elegant_swirles/processing/internal/config"
	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
	"github.com/defaulterrr/elegant_swirles/processing/internal/service"
	"github.com/defaulterrr/elegant_swirles/processing/metrics"
)

type Server struct {
	Service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{
		Service: service,
	}
}

func (s *Server) Start(cfg *config.Metrics) error {
	metricsServer := createMetricsServer(cfg)

	go func() {
		fmt.Printf("Metrics server is running on %s:%d\n", cfg.Host, cfg.Port)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Failed running metrics server:%v\n", err)
		}
	}()

	return nil
}

func (s *Server) ReadDHTMetrics(ctx context.Context) error {
	curMetrics := make(chan model.DHTMetrics)

	go func() {
		err := s.Service.DHTService.GetDHTMetrics(ctx, curMetrics)
		if err != nil {
			fmt.Printf("GetDHTMetrics: %v\n", err)
		}
	}()

	for metr := range curMetrics {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		metrics.SetTemperature(float64(metr.Temperature))
		metrics.SetHumidity(float64(metr.Humidity))
	}

	return nil
}

func (s *Server) ReadCameraMetrics(ctx context.Context) error {
	curCameraMetrics := make(chan model.CameraMetrics)

	go func() {
		err := s.Service.CameraService.GetCameraMetrics(ctx, curCameraMetrics)
		if err != nil {
			fmt.Printf("GetCameraMetrics: %v\n", err)
		}
	}()

	for cameraMetr := range curCameraMetrics {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		metrics.SetCountPeople(float64(cameraMetr.CountPeople))
	}

	return nil
}
