package server

import (
	"context"
	"fmt"
	"net"

	"github.com/defaulterrr/elegant_swirles/camera/grpc/go/elegant_swirles/camera"
	"github.com/defaulterrr/elegant_swirles/camera/internal/config"
	"github.com/defaulterrr/elegant_swirles/camera/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ICameraService interface {
	GetMetrics(ctx context.Context, metrics chan<- model.CameraMetrics) error
}

type Server struct {
	cameraService ICameraService
	camera.UnimplementedCameraServer
}

func NewServer(cameraService ICameraService) *Server {
	return &Server{
		cameraService: cameraService,
	}
}

func (s *Server) Start(grpcCfg *config.Grpc) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%v", grpcCfg.Host, grpcCfg.Port))
	if err != nil {
		return fmt.Errorf("Listen: %v", err)
	}

	newServer := grpc.NewServer()
	camera.RegisterCameraServer(newServer, s)
	if err := newServer.Serve(lis); err != nil {
		return fmt.Errorf("Serve: %v", err)
	}

	return nil
}

func (s *Server) GetCameraMetrics(in *emptypb.Empty, cameraServer camera.Camera_GetCameraMetricsServer) error {
	metrics := make(chan model.CameraMetrics)

	go func() {
		err := s.cameraService.GetMetrics(cameraServer.Context(), metrics)
		if err != nil {
			fmt.Printf("GetMetrics: %v\n", err)
		}
	}()

	for el := range metrics {
		if err := cameraServer.Send(&camera.CameraMetrics{
			CountPeople: el.CountPeople,
			Created:     timestamppb.Now(),
		}); err != nil {
			fmt.Printf("Send: %v\n", err)
			return err
		}
	}

	return nil
}
