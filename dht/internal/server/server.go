package server

import (
	"fmt"
	"net"

	pb "github.com/defaulterrr/iot3/dht/grpc/go/um4ru_ch4n/dht"
	"github.com/defaulterrr/iot3/dht/internal/config"
	"github.com/defaulterrr/iot3/dht/internal/model"
	"github.com/defaulterrr/iot3/dht/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	serv service.IService
	pb.UnimplementedDHTServer
}

func NewServer(serv service.IService) *Server {
	return &Server{
		serv: serv,
	}
}

func (s *Server) Start(cfg *config.Grpc) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%v", cfg.Host, cfg.Port))
	if err != nil {
		return fmt.Errorf("Listen: %v", err)
	}

	newServer := grpc.NewServer()
	pb.RegisterDHTServer(newServer, s)
	if err := newServer.Serve(lis); err != nil {
		return fmt.Errorf("Serve: %v", err)
	}

	return nil
}

func (s *Server) GetDHTMetrics(in *emptypb.Empty, srv pb.DHT_GetDHTMetricsServer) error {
	metrics := make(chan model.DHTMetrics)

	go func() {
		err := s.serv.GetMetrics(srv.Context(), metrics)
		if err != nil {
			fmt.Printf("GetMetrics: %v\n", err)
		}
	}()

	for el := range metrics {
		if err := srv.Send(&pb.Metrics{
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
