package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"math/rand"
	"ozon/internal/links"
	"ozon/internal/models"
	"ozon/pkg/httpErrors"
	"time"
)

type linksRepo struct {
	db *sqlx.DB
}

func NewLinksRepository(db *sqlx.DB) links.Repository {
	return &linksRepo{db: db}
}

func (r *linksRepo) Create(ctx context.Context, links *models.Links) (*models.LinkShort, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Register")
	defer span.Finish()

	rm := &models.LinkShort{}

	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "_"

	var shortLink string
	for i := 0; i < 10; i++ {
		shortLink += string(charset[rand.Intn(len(charset))])
	}

	links.ShortLink = shortLink

	fmt.Println(shortLink)

	if err := r.db.QueryRowxContext(ctx, createLink, &links.Link, &links.ShortLink).StructScan(rm); err != nil {
		println(httpErrors.NewExistsLinkError("link Already Exists"))
		return nil, errors.Wrap(err, "postRepo.Register.StructScan")
	}

	return rm, nil
}

func (r *linksRepo) Read(ctx context.Context, shortLink string) (*models.ReadLinks, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentsRepo.Read")
	defer span.Finish()

	rows, err := r.db.QueryxContext(ctx, readLink, shortLink)
	if err != nil {
		return nil, httpErrors.NewNotFoundError("Short link not found")
	}
	defer rows.Close()

	var originalLink string
	found := false
	for rows.Next() {
		found = true
		if err = rows.Scan(&originalLink); err != nil {
			return nil, errors.Wrap(err, "linksRepo.Read.Scan")
		}
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "linksRepo.Read.Rows.Err")
	}

	if !found {
		return nil, httpErrors.NewNotFoundError("Short link not found")
	}

	return &models.ReadLinks{
		Link: originalLink,
	}, nil
}
