package db

import (
	"fmt"

	"github.com/PentaGol/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	connDb, err := sqlx.Open("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDb, nil
}
