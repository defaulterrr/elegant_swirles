package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/defaulterrr/elegant_swirles/processing/internal/config"
	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
)

type IDHTService interface {
	GetDHTMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error
}

type Server struct {
	dhtService IDHTService
}

func NewServer(dhtService IDHTService) *Server {
	return &Server{
		dhtService: dhtService,
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
