package adapters

import (
	"butify/pkg/models"
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type PqUrls struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *PqUrls {
	return &PqUrls{db: db}
}

func (pu *PqUrls) CreateUrl(ctx context.Context, url models.Url) (models.Url, error) {
	query := "INSERT INTO urls (small, origin, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	if url.CreatedAt.IsZero() {
		url.CreatedAt = time.Now().UTC()
	}
	if url.UpdateAt.IsZero() {
		url.UpdateAt = time.Now().UTC()
	}

	err := pu.db.QueryRowContext(ctx, query, url.Small, url.Origin, url.CreatedAt, url.UpdateAt).Scan(&url.Id)

	return url, err
}

func (pu *PqUrls) GetBySmall(ctx context.Context, small string) (models.Url, error) {
	query := "SELECT id, small, origin, created_at, updated_at FROM urls WHERE small=$1"

	url := models.Url{}
	err := pu.db.Get(&url, query, small)
	return url, err
}
