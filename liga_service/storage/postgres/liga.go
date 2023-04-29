package postgres

import (
	"log"
	"time"

	p "github.com/PentaGol/liga_service/genproto/liga"
)

func (r *LigaRepo) CreateLiga(liga *p.LigaRequest) (*p.LigaResponse, error) {
	var res p.LigaResponse
	err := r.db.QueryRow(`
		insert into 
			posts(title, description, admin_id) 
		values
			($1, $2, $3) 
		returning 
			id, title, description, admin_id, created_at, updated_at`, liga.Title, liga.Description, liga.AdminId).
		Scan(
			&res.Id,
			&res.Title,
			&res.Description,
			&res.AdminId,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to create liga")
		return &p.LigaResponse{}, err
	}

	return &res, nil
}

func (r *LigaRepo) GetLigaById(liga *p.IdRequest) (*p.LigaResponse, error) {
	res := p.LigaResponse{}
	err := r.db.QueryRow(`
		select 
			id, title, description, admin_id, created_at, updated_at 
		from 
			posts 
		where 
			id = $1 and deleted_at is null`, liga.Id).
		Scan(
			&res.Id,
			&res.Title,
			&res.Description,
			&res.AdminId,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to get post")
		return &p.LigaResponse{}, err
	}

	return &res, nil
}

func (r *LigaRepo) GetAllLigas(req *p.AllLigaRequest) (*p.Ligas, error) {
	res := p.Ligas{}

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Query(`
		select 
			id, title, description, admin_id, created_at, updated_at 
		from 
			ligas 
		where 
			deleted_at is null 
		limit $1 offset $2`, req.Limit, offset,
	)
	if err != nil {
		log.Println("failed to get post")
		return &p.Ligas{}, err
	}

	for rows.Next() {
		temp := p.LigaResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.Title,
			&temp.Description,
			&temp.AdminId,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			log.Println("failed to scanning post")
			return &p.Ligas{}, err
		}

		res.Ligas = append(res.Ligas, &temp)
	}

	return &res, nil
}

func (r *LigaRepo) DeleteLiga(id *p.IdRequest) (*p.LigaResponse, error) {
	post := p.LigaResponse{}
	err := r.db.QueryRow(`
		update 
			ligas 
		set 
			deleted_at = $1 
		where 
			id = $2 
		returning 
			id, title, description, admin_id, created_at, updated_at`, time.Now(), id.Id).
		Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.AdminId,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to delete post")
		return &p.LigaResponse{}, err
	}

	return &post, nil
}
