package service

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"ozon/config"
	"ozon/internal/links"
	"ozon/internal/models"
)

type linksService struct {
	cfg       *config.Config
	linksRepo links.Repository
}

func NewLinksService(cfg *config.Config, linksRepo links.Repository) links.Service {
	return &linksService{cfg: cfg, linksRepo: linksRepo}
}

func (r *linksService) Create(ctx context.Context, links *models.Links) (*models.LinkShort, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "linksUC.Create")
	defer span.Finish()
	rm, err := r.linksRepo.Create(ctx, links)
	if err != nil {
		fmt.Println("linksRepo.Create", err)
		return nil, err
	}

	return rm, err
}

func (r *linksService) Read(ctx context.Context, shortLink string) (*models.ReadLinks, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "linksUC.Read")
	defer span.Finish()

	return r.linksRepo.Read(ctx, shortLink)
}
