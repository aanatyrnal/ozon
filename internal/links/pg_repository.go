package links

import (
	"context"
	"ozon/internal/models"
)

type Repository interface {
	Create(ctx context.Context, link *models.Links) (*models.LinkShort, error)
	Read(ctx context.Context, shortLink string) (*models.ReadLinks, error)
}
