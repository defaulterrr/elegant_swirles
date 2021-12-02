package repository

import (
	"context"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/grpc/go/elegant_swirles/dht"
	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DHTMetricsGRPC struct {
	DHTConn *grpc.ClientConn
}

func NewDHTMetricsGRPC(DHTConn *grpc.ClientConn) *DHTMetricsGRPC {
	return &DHTMetricsGRPC{
		DHTConn: DHTConn,
	}
}

func (m *DHTMetricsGRPC) GetFromDHT(ctx context.Context, metrics chan<- model.DHTMetrics) error {
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

func (m *DHTMetricsGRPC) getDHTClient() dht.DHTClient {
	return dht.NewDHTClient(m.DHTConn)
}

func dhtMetricsPbToModel(metrics *dht.Metrics) model.DHTMetrics {
	return model.DHTMetrics{
		Temperature: metrics.GetTemperature(),
		Humidity:    metrics.GetHumidity(),
		Created:     metrics.GetCreated().Seconds,
	}
}
