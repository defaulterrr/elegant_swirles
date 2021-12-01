package service

import (
	"context"
	"fmt"

	"github.com/defaulterrr/elegant_swirles/processing/internal/model"
	"github.com/defaulterrr/elegant_swirles/processing/internal/repository"
)

type IService interface {
	GetDHTMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error
}

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) IService {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) GetDHTMetrics(ctx context.Context, metrics chan<- model.DHTMetrics) error {
	err := s.Repo.GetFromDHT(ctx, metrics)
	if err != nil {
		return fmt.Errorf("GetFromDHT: %v", err)
	}

	return nil
}
