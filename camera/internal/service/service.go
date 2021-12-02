package service

import (
	"context"
	"math"
	"time"

	"github.com/defaulterrr/elegant_swirles/camera/internal/model"
)

type CameraService struct {
}

func NewService() *CameraService {
	return &CameraService{}
}

func (s *CameraService) GetMetrics(ctx context.Context, metrics chan<- model.CameraMetrics) error {
	var i float64

	for {
		newMetrics := model.CameraMetrics{
			CountPeople: uint32(math.Round(math.Abs(math.Sin(i) * 20))),
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
