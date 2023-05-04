package repo

import (
	p "github.com/PentaGol/genproto/liga"
)

type ClubStorageI interface {
	CreateClub(*p.ClubRequest) (*p.ClubResponse, error)
	GetClubById(*p.IdRequest) (*p.ClubResponse, error)
	GetAllClubs(*p.AllClubRequest) (*p.Clubs, error)
}
