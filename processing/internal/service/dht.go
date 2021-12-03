package service

import (
	"context"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
)

type DHTRepo interface {
	GetFromDHT(ctx context.Context, metrics chan<- model.DHTMetrics) error
}

type DHTService struct {
	DHTRepo DHTRepo
}

func NewDHTService(dhtRepo DHTRepo) *DHTService {
	return &DHTService{
		DHTRepo: dhtRepo,
	}
}

func (s *DHTService) GetDHTMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error {
	err := s.DHTRepo.GetFromDHT(ctx, metrics)
	if err != nil {
		return fmt.Errorf("GetFromDHT: %v", err)
	}

	return nil
}
