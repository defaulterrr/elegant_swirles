package server

import (
	"context"
	"fmt"
	"net"

	"github.com/defaulterrr/elegant_swirles/dht/grpc/go/elegant_swirles/dht"
	"github.com/defaulterrr/elegant_swirles/dht/internal/config"
	"github.com/defaulterrr/elegant_swirles/dht/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IDHTService interface {
	GetMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error
}

type Server struct {
	dhtService IDHTService
	dht.UnimplementedDHTServer
}

func NewServer(dhtService IDHTService) *Server {
	return &Server{
		dhtService: dhtService,
	}
}

func (s *Server) Start(grpcCfg *config.Grpc) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%v", grpcCfg.Host, grpcCfg.Port))
	if err != nil {
		return fmt.Errorf("Listen: %v", err)
	}

	newServer := grpc.NewServer()
	dht.RegisterDHTServer(newServer, s)
	if err := newServer.Serve(lis); err != nil {
		return fmt.Errorf("Serve: %v", err)
	}

	return nil
}

func (s *Server) GetDHTMetrics(in *emptypb.Empty, dhtServer dht.DHT_GetDHTMetricsServer) error {
	metrics := make(chan model.DHTMetrics)

	go func() {
		err := s.dhtService.GetMetrics(dhtServer.Context(), metrics)
		if err != nil {
			fmt.Printf("GetMetrics: %v\n", err)
		}
	}()

	for el := range metrics {
		if err := dhtServer.Send(&dht.Metrics{
			Temperature: el.Temperature,
			Humidity:    el.Temperature,
			Created:     timestamppb.Now(),
		}); err != nil {
			fmt.Printf("Send: %v\n", err)
			return err
		}
	}

	return nil
}
