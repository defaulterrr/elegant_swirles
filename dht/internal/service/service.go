package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/defaulterrr/elegant_swirles/dht/internal/model"
)

type DHTService struct {
}

func NewService() *DHTService {
	return &DHTService{}
}

func (s *DHTService) GetMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error {
	rand.Seed(time.Now().UnixNano())

	for {
		newMetrics := model.DHTMetrics{
			Temperature: float32(rand.Int31n(50)),
			Humidity:    float32(rand.Int31n(100)),
		}

		select {
		case <-ctx.Done():
			close(metrics)
			return ctx.Err()
		case metrics <- newMetrics:
		}
		time.Sleep(time.Second * 1)
	}
}
