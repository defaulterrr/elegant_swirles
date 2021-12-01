package repository

type Repository struct {
	Metrics
}

func NewRepository(conns *GRPCConns) *Repository {
	return &Repository{
		Metrics: NewMetricsGRPC(conns.DHTConn),
	}
}
