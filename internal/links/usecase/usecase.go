package usecase

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"ozon/config"
	"ozon/internal/links"
	"ozon/internal/models"
)

type linksUC struct {
	cfg          *config.Config
	linksService links.Service
}

func NewLinksUseCase(cfg *config.Config, linksService links.Service) links.UseCase {
	return &linksUC{cfg: cfg, linksService: linksService}
}

func (r *linksUC) Create(ctx context.Context, links *models.Links) (*models.LinkShort, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "linksUC.Create")
	defer span.Finish()

	return r.linksService.Create(ctx, links)
}

func (r *linksUC) Read(ctx context.Context, shortLink string) (*models.ReadLinks, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "linksUC.Read")
	defer span.Finish()

	return r.linksService.Read(ctx, shortLink)
}
