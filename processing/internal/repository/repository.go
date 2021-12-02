package repository

import (
	"context"

	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
)

type DHTMetrics interface {
	GetFromDHT(ctx context.Context, metrics chan<- model.DHTMetrics) error
}

type ICameraMetrics interface {
	GetFromCamera(ctx context.Context, metrics chan<- model.CameraMetrics) error
}

type Repository struct {
	DHTMetrics    DHTMetrics
	CameraMetrics ICameraMetrics
}

func NewRepository(conns *GRPCConns) *Repository {
	return &Repository{
		DHTMetrics:    NewDHTMetricsGRPC(conns.DHTConn),
		CameraMetrics: NewCameraMetricsGRPC(conns.CameraConn),
	}
}
