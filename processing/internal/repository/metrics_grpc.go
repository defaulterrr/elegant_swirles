package repository

import (
	"context"
	"fmt"

	pb "github.com/defaulterrr/elegant_swirles/processing/grpc/go/um4ru_ch4n/dht"
	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Metrics interface {
	GetFromDHT(ctx context.Context, metrics chan<- model.DHTMetrics) error
}

type MetricsGRPC struct {
	DHTConn *grpc.ClientConn
}

func NewMetricsGRPC(DHTConn *grpc.ClientConn) Metrics {
	return &MetricsGRPC{
		DHTConn: DHTConn,
	}
}

func (m *MetricsGRPC) GetFromDHT(ctx context.Context, metrics chan<- model.DHTMetrics) error {
	stream, err := m.getDHTClient().GetDHTMetrics(ctx, &emptypb.Empty{})
	if err != nil {
		close(metrics)
		return fmt.Errorf("GetDHTMetrics: %v", err)
	}

	for {
		metr, err := stream.Recv()
		if err != nil {
			close(metrics)
			return err
		}

		select {
		case <-ctx.Done():
			close(metrics)
			return ctx.Err()
		case metrics <- dhtMetricsPbToModel(metr):
		}
	}
}

func (m *MetricsGRPC) getDHTClient() pb.DHTClient {
	return pb.NewDHTClient(m.DHTConn)
}

func dhtMetricsPbToModel(metr *pb.Metrics) model.DHTMetrics {
	return model.DHTMetrics{
		Temperature: metr.GetTemperature(),
		Humidity:    metr.GetHumidity(),
		Created:     metr.GetCreated().Seconds,
	}
}
