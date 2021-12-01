package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	temperature prometheus.Gauge
	humidity    prometheus.Gauge
)

func InitMetrics() error {
	temperature = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temperature",
		Help: "Current temperature from DHT",
	})

	humidity = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "humidity",
		Help: "Current humidity from DHT",
	})

	err := prometheus.Register(temperature)
	if err != nil {
		return fmt.Errorf("prometheus.Register(temperature): %v", err)
	}

	err = prometheus.Register(humidity)
	if err != nil {
		return fmt.Errorf("prometheus.Register(humidity): %v", err)
	}

	return nil
}

func SetTemperature(temp float64) {
	if temperature == nil {
		return
	}

	temperature.Set(temp)
}

func SetHumidity(hum float64) {
	if humidity == nil {
		return
	}

	humidity.Set(hum)
}
