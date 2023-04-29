package postgres

import "github.com/jmoiron/sqlx"

type LigaRepo struct {
	db *sqlx.DB
}

func NewLigaRepo(db *sqlx.DB) *LigaRepo {
	return &LigaRepo{
		db: db,
	}
}
