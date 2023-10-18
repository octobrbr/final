package storage

import (
	"censor/pkg/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(connstr string) (*DB, error) {

	if connstr == "" {
		return nil, errors.New("wrong connection string")
	}

	pool, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}
	db := DB{
		pool: pool,
	}
	return &db, nil
}

func (db *DB) GetBlackList() ([]models.BannedWord, error) {
	rows, err := db.pool.Query(context.Background(), "SELECT * FROM blacklist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.BannedWord
	for rows.Next() {
		var c models.BannedWord
		err = rows.Scan(&c.ID, &c.Word)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, rows.Err()
}
