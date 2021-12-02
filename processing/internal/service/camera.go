package service

import (
	"context"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
)

type CameraRepo interface {
	GetFromCamera(ctx context.Context, metrics chan<- model.CameraMetrics) error
}

type CameraService struct {
	CameraRepo CameraRepo
}

func NewCameraService(cameraRepo CameraRepo) *CameraService {
	return &CameraService{
		CameraRepo: cameraRepo,
	}
}

func (c *CameraService) GetCameraMetrics(ctx context.Context, metrics chan<- model.CameraMetrics) error {
	err := c.CameraRepo.GetFromCamera(ctx, metrics)
	if err != nil {
		return fmt.Errorf("GetFromCamera: %v", err)
	}

	return nil
}
