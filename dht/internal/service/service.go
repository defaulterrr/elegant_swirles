package service

import (
	"context"
	"math"
	"time"

	"github.com/defaulterrr/elegant_swirles/dht/internal/model"
)

type DHTService struct {
}

func NewService() *DHTService {
	return &DHTService{}
}

func (s *DHTService) GetMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error {
	var i float64

	for {
		newMetrics := model.DHTMetrics{
			Temperature: float32(math.Abs(math.Sin(i) * 50)),
			Humidity:    float32(math.Abs(math.Cos(i) * 100)),
		}

		select {
		case <-ctx.Done():
			close(metrics)
			return ctx.Err()
		case metrics <- newMetrics:
		}

		i = float64(int(i+1) % 1000)
		time.Sleep(time.Second * 1)
	}
}
