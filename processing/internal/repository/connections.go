package repository

import (
	"fmt"

	"github.com/defaulterrr/iot3/processing/internal/config"
	"google.golang.org/grpc"
)

type GRPCConns struct {
	DHTConn *grpc.ClientConn
}

func GetGRPCConns(cfg *config.Grpc) (*GRPCConns, error) {
	dhtConn, err := getDHTConn(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("getDHTConn: %v", err)
	}

	return &GRPCConns{
		DHTConn: dhtConn,
	}, nil
}

func getDHTConn(target string) (*grpc.ClientConn, error) {
	return grpc.Dial(target, grpc.WithInsecure())
}
