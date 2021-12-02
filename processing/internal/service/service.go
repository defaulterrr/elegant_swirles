package service

import (
	"context"

	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
	"github.com/defaulterrr/elegant_swirles/processing/internal/repository"
)

type IDHTService interface {
	GetDHTMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error
}

type ICameraService interface {
	GetCameraMetrics(ctx context.Context, metrics chan<- model.CameraMetrics) error
}

type Service struct {
	DHTService    IDHTService
	CameraService ICameraService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DHTService:    NewDHTService(repos.DHTMetrics),
		CameraService: NewCameraService(repos.CameraMetrics),
	}
}
