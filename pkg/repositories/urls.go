package repositories

import (
	"butify/pkg/adapters"
	"butify/pkg/models"
	"context"
)

type Urls struct {
	adapter *adapters.PqUrls
}

func New(adapter *adapters.PqUrls) *Urls {
	return &Urls{adapter: adapter}
}

func (u *Urls) CreateUrl(ctx context.Context, url models.Url) (models.Url, error) {
	return u.adapter.CreateUrl(ctx, url)
}

func (u *Urls) GetBySmall(ctx context.Context, small string) (models.Url, error) {
	return u.adapter.GetBySmall(ctx, small)
}

func (u *Urls) GenerateUrl(ctx context.Context) string {
	return "small"
}

func (u *Urls) CheckUrlIsValid(ctx context.Context, url string) bool {
	return true
}
