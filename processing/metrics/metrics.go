package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var temperature prometheus.Gauge

func InitMetrics() error {
	temperature = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temperature",
		Help: "Current temperature",
	})

	err := prometheus.Register(temperature)
	if err != nil {
		return fmt.Errorf("prometheus.Register(temperature): %v", err)
	}

	return nil
}

func SetTemperature(temp float64) {
	if temperature == nil {
		return
	}

	temperature.Set(temp)
}
