package postgres

import (
	"log"
	"time"

	p "github.com/PentaGol/liga_service/genproto/liga"
)

func (r *Repo) CreateGame(game *p.GameRequest) (*p.GameResponse, error) {
	var res p.GameResponse
	err := r.db.QueryRow(`
		insert into 
			games(name) 
		values
			($1) 
		returning 
			id, name, created_at, updated_at`, game.Time).
		Scan(
			&res.Id,
			&res.Time,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to create liga")
		return &p.GameResponse{}, err
	}

	return &res, nil
}

func (r *Repo) GetGameById(game *p.IdRequest) (*p.GameResponse, error) {
	res := p.GameResponse{}
	err := r.db.QueryRow(`
		select 
			id, name, created_at, updated_at 
		from 
			games 
		where 
			id = $1 and deleted_at is null`, game.Id).
		Scan(
			&res.Id,
			&res.Time,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to get liga")
		return &p.GameResponse{}, err
	}

	return &res, nil
}

func (r *Repo) GetAllGames(req *p.AllGameRequest) (*p.Games, error) {
	res := p.Games{}

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Query(`
		select 
			id, name, created_at, updated_at 
		from 
			games 
		where 
			deleted_at is null 
		limit $1 offset $2`, req.Limit, offset,
	)
	if err != nil {
		log.Println("failed to get game")
		return &p.Games{}, err
	}

	for rows.Next() {
		temp := p.GameResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.Time,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			log.Println("failed to scanning liga")
			return &p.Games{}, err
		}

		res.Games = append(res.Games, &temp)
	}

	return &res, nil
}

func (r *Repo) DeleteGame(id *p.IdRequest) (*p.GameResponse, error) {
	game := p.GameResponse{}
	err := r.db.QueryRow(`
		update 
			games 
		set 
			deleted_at = $1 
		where 
			id = $2 and deleted_at is null
		returning 
			id, name, created_at, updated_at`, time.Now(), id.Id).
		Scan(
			&game.Id,
			&game.Time,
			&game.CreatedAt,
			&game.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to delete liga")
		return &p.GameResponse{}, err
	}

	return &game, nil
}
