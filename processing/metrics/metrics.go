package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	temperature prometheus.Gauge
	humidity    prometheus.Gauge
	countPeople prometheus.Gauge
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

	countPeople = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "countPeople",
		Help: "Current countPeople from Camera",
	})

	err := prometheus.Register(temperature)
	if err != nil {
		return fmt.Errorf("prometheus.Register(temperature): %v", err)
	}

	err = prometheus.Register(humidity)
	if err != nil {
		return fmt.Errorf("prometheus.Register(humidity): %v", err)
	}

	err = prometheus.Register(countPeople)
	if err != nil {
		return fmt.Errorf("prometheus.Register(countPeople): %v", err)
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

func SetCountPeople(count float64) {
	if countPeople == nil {
		return
	}

	countPeople.Set(count)
}
