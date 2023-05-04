package repo

import (
	p "github.com/PentaGol/genproto/liga"
)

type GameStorageI interface {
	CreateGame(*p.GameRequest) (*p.GameResponse, error)
	GetGameById(*p.IdRequest) (*p.GameResponse, error)
	GetAllGames(*p.AllGameRequest) (*p.Games, error)
	DeleteGame(*p.IdRequest) (*p.GameResponse, error)
}
