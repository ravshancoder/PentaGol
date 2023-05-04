package storage

import (
	"github.com/PentaGol/storage/postgres"
	"github.com/PentaGol/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Admin() repo.AdminStoreI
	Liga() repo.LigaStorageI
	Club() repo.ClubStorageI
	Game() repo.GameStorageI
	Post() repo.PostStorageI
}

type storagePg struct {
	db        *sqlx.DB
	adminRepo repo.AdminStoreI
	postRepo  repo.PostStorageI
	ligaRepo  repo.LigaStorageI
	gameRepo  repo.GameStorageI
	clubRepo  repo.ClubStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:        db,
		adminRepo: postgres.NewRepo(db),
		postRepo:  postgres.NewRepo(db),
		ligaRepo:  postgres.NewRepo(db),
		gameRepo:  postgres.NewRepo(db),
		clubRepo:  postgres.NewRepo(db),
	}
}

func (s storagePg) Admin() repo.AdminStoreI {
	return s.adminRepo
}

func (s storagePg) Post() repo.PostStorageI {
	return s.postRepo
}

func (s storagePg) Liga() repo.LigaStorageI {
	return s.ligaRepo
}

func (s storagePg) Game() repo.GameStorageI {
	return s.gameRepo
}

func (s storagePg) Club() repo.ClubStorageI {
	return s.clubRepo
}
