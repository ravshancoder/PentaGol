package repo

import (
	p "github.com/PentaGol/liga_service/genproto/liga"
)

type LigaStorageI interface {
	CreateLiga(*p.LigaRequest) (*p.LigaResponse, error)
	GetLigaById(*p.IdRequest) (*p.LigaResponse, error)
	GetAllLigas(*p.AllLigaRequest) (*p.Ligas, error)
	DeleteLiga(*p.IdRequest) (*p.LigaResponse, error)
}
