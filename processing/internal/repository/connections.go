package repository

import (
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/internal/config"
	"google.golang.org/grpc"
)

type GRPCConns struct {
	DHTConn    *grpc.ClientConn
	CameraConn *grpc.ClientConn
}

func GetGRPCConns(cfg *config.Config) (*GRPCConns, error) {
	dhtConn, err := getDHTConn(fmt.Sprintf("%s:%d", cfg.DHTGrpc.Host, cfg.DHTGrpc.Port))
	if err != nil {
		return nil, fmt.Errorf("getDHTConn: %v", err)
	}

	cameraConn, err := getCameraConn(fmt.Sprintf("%s:%d", cfg.CameraGrpc.Host, cfg.CameraGrpc.Port))
	if err != nil {
		return nil, fmt.Errorf("getCameraConn: %v", err)
	}

	return &GRPCConns{
		DHTConn:    dhtConn,
		CameraConn: cameraConn,
	}, nil
}

func getDHTConn(target string) (*grpc.ClientConn, error) {
	return grpc.Dial(target, grpc.WithInsecure())
}

func getCameraConn(target string) (*grpc.ClientConn, error) {
	return grpc.Dial(target, grpc.WithInsecure())
}
