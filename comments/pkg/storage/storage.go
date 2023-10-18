package storage

import (
	"comments/pkg/models"
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

func (db *DB) GetAllComments(newsID int) ([]models.Comment, error) {
	rows, err := db.pool.Query(context.Background(), "SELECT * FROM comments WHERE news_id = $1;", newsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		err = rows.Scan(&c.ID, &c.NewsID, &c.Content, &c.PubTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, rows.Err()
}

func (db *DB) AddComment(c models.Comment) error {
	_, err := db.pool.Exec(context.Background(),
		"INSERT INTO comments (news_id,content) VALUES ($1,$2);", c.NewsID, c.Content)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteComment(c models.Comment) error {
	_, err := db.pool.Exec(context.Background(),
		"DELETE FROM comments WHERE id=$1;", c.ID)
	if err != nil {
		return err
	}
	return nil
}
