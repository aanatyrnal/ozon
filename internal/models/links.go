package models

import "github.com/jmoiron/sqlx"

type Links struct {
	Id        int    `json:"id,omitempty" db:"id"`
	Link      string `json:"link,omitempty" db:"link"`
	ShortLink string `json:"short_link,omitempty" db:"short_link"`
}

type LinkShort struct {
	ShortLink string `json:"short_link,omitempty" db:"short_link"`
}

type CreateLinks struct {
	Id   int    `json:"id,omitempty" db:"id"`
	Link string `json:"link,omitempty" db:"link"`
}

type ReadLinks struct {
	Link string `json:"link,omitempty" db:"link"`
}

type linksRepo struct {
	db *sqlx.DB
}
