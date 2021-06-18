package services

import (
	"butify/pkg/models"
	"context"
)

type Repository interface {
	CreateUrl(ctx context.Context, url models.Url) (models.Url, error)
	GetBySmall(ctx context.Context, small string) (models.Url, error)
	GenerateUrl(ctx context.Context) string
	CheckUrlIsValid(ctx context.Context, url string) bool
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUrl(ctx context.Context, url models.Url) (models.Url, error) {
	return url, nil
}

func (s *Service) GetBySmallUrl(ctx context.Context, small string) (models.Url, error) {
	return s.repo.GetBySmall(ctx, small)
}
