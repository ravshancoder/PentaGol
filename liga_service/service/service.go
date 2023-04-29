package service

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"

	p "github.com/PentaGol/liga_service/genproto/liga"
	"github.com/PentaGol/liga_service/pkg/logger"
	grpcclient "github.com/PentaGol/liga_service/service/grpc_client"
	"github.com/PentaGol/liga_service/storage"
)

type LigaService struct {
	storage storage.IStorage
	Logger  logger.Logger
	Client  grpcclient.Clients
}

func NewLigaService(db *sqlx.DB, log logger.Logger, client grpcclient.Clients) *LigaService {
	return &LigaService{
		storage: storage.NewStoragePg(db),
		Logger:  log,
		Client:  client,
	}
}

func (s *LigaService) CreateLiga(ctx context.Context, req *p.LigaRequest) (*p.LigaResponse, error) {
	res, err := s.storage.Liga().CreateLiga(req)
	if err != nil {
		log.Println("failed to create Liga: ", err)
		return &p.LigaResponse{}, err
	}

	return res, nil
}

func (s *LigaService) GetLigaById(ctx context.Context, req *p.IdRequest) (*p.LigaResponse, error) {
	res, err := s.storage.Liga().GetLigaById(req)
	if err != nil {
		log.Println("failed to get Liga by id: ", err)
		return &p.LigaResponse{}, err
	}

	return res, nil
}

func (s *LigaService) GetAllLigas(ctx context.Context, req *p.AllLigaRequest) (*p.Ligas, error) {
	res, err := s.storage.Liga().GetAllLigas(req)
	if err != nil {
		log.Println("failed to get all Liga: ", err)
		return &p.Ligas{}, err
	}

	return res, nil
}

func (s *LigaService) DeleteLiga(ctx context.Context, req *p.IdRequest) (*p.LigaResponse, error) {
	res, err := s.storage.Liga().DeleteLiga(req)
	if err != nil {
		log.Println("failed to delete Liga: ", err)
		return &p.LigaResponse{}, err
	}

	return res, err
}
