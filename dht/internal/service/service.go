package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/defaulterrr/elegant_swirles/dht/internal/model"
)

type IService interface {
	GetMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error
}

type Service struct {
}

func NewService() IService {
	return &Service{}
}

func (s *Service) GetMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error {
	rand.Seed(time.Now().UnixNano())

	for {
		newMetr := model.DHTMetrics{
			Temperature: float32(rand.Int31n(50)),
			Humidity:    float32(rand.Int31n(100)),
		}

		select {
		case <-ctx.Done():
			close(metrics)
			return ctx.Err()
		case metrics <- newMetr:
		}
		time.Sleep(time.Second * 1)
	}
}
