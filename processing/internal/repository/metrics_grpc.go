package repository

import (
	"context"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/grpc/go/um4ru_ch4n/dht"
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
	metricsStream, err := m.getDHTClient().GetDHTMetrics(ctx, &emptypb.Empty{})
	if err != nil {
		close(metrics)
		return fmt.Errorf("GetDHTMetrics: %v", err)
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
		case metrics <- dhtMetricsPbToModel(curMetrics):
		}
	}
}

func (m *MetricsGRPC) getDHTClient() dht.DHTClient {
	return dht.NewDHTClient(m.DHTConn)
}

func dhtMetricsPbToModel(metrics *dht.Metrics) model.DHTMetrics {
	return model.DHTMetrics{
		Temperature: metrics.GetTemperature(),
		Humidity:    metrics.GetHumidity(),
		Created:     metrics.GetCreated().Seconds,
	}
}
