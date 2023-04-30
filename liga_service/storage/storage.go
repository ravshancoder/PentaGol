package storage

import (
	"github.com/PentaGol/liga_service/storage/postgres"
	"github.com/PentaGol/liga_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Liga() repo.LigaStorageI
	Game() repo.GameStorageI
}

type storagePg struct {
	db       *sqlx.DB
	ligaRepo repo.LigaStorageI
	gameRepo repo.GameStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		ligaRepo: postgres.NewLigaRepo(db),
		gameRepo: postgres.NewGameRepo(db),
	}
}

func (s storagePg) Liga() repo.LigaStorageI {
	return s.ligaRepo
}


func (s storagePg) Game() repo.GameStorageI {
	return s.gameRepo
}