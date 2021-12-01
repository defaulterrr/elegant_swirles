package service

import (
	"context"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
	"github.com/defaulterrr/elegant_swirles/processing/internal/repository"
)

type DHTService struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *DHTService {
	return &DHTService{
		repo: repo,
	}
}

func (s *DHTService) GetDHTMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error {
	err := s.repo.GetFromDHT(ctx, metrics)
	if err != nil {
		return fmt.Errorf("GetFromDHT: %v", err)
	}

	return nil
}
