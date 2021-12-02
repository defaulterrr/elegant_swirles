package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/defaulterrr/elegant_swirles/processing/internal/config"
	"github.com/defaulterrr/elegant_swirles/processing/internal/service"
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
