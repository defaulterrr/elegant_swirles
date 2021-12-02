package repository

import (
	"context"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/grpc/go/elegant_swirles/camera"
	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CameraMetricsGRPC struct {
	CameraConn *grpc.ClientConn
}

func NewCameraMetricsGRPC(cameraConn *grpc.ClientConn) *CameraMetricsGRPC {
	return &CameraMetricsGRPC{
		CameraConn: cameraConn,
	}
}

func (c *CameraMetricsGRPC) GetFromCamera(ctx context.Context, metrics chan<- model.CameraMetrics) error {
	metricsStream, err := c.getCameraClient().GetCameraMetrics(ctx, &emptypb.Empty{})
	if err != nil {
		close(metrics)
		return fmt.Errorf("GetCameraMetrics: %v", err)
	}

	for {
		curMetrics, err := metricsStream.Recv()
		if err != nil {
			close(metrics)
			return err
		}

		select {
		case <-ctx.Done():
			close(metrics)
			return ctx.Err()
		case metrics <- cameraMetricsPbToModel(curMetrics):
		}
	}
}

func (c *CameraMetricsGRPC) getCameraClient() camera.CameraClient {
	return camera.NewCameraClient(c.CameraConn)
}

func cameraMetricsPbToModel(metrics *camera.CameraMetrics) model.CameraMetrics {
	return model.CameraMetrics{
		CountPeople: metrics.GetCountPeople(),
		Created:     metrics.GetCreated().Seconds,
	}
}
