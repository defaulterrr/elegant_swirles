package server

import (
	"fmt"
	"net/http"

	"github.com/defaulterrr/elegant_swirles/processing/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func CreateMetricsServer(cfg *config.Metrics) *http.Server {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	mux := http.DefaultServeMux
	mux.Handle(cfg.Path, promhttp.Handler())

	metricsServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return metricsServer
}
