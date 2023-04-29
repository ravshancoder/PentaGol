package storage

import (
	"github.com/PentaGol/liga_service/storage/postgres"
	"github.com/PentaGol/liga_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Liga() repo.LigaStorageI
}

type storagePg struct {
	db       *sqlx.DB
	ligaRepo repo.LigaStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		ligaRepo: postgres.NewLigaRepo(db),
	}
}

func (s storagePg) Liga() repo.LigaStorageI {
	return s.ligaRepo
}
