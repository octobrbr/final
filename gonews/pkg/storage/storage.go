package storage

import (
	"context"
	"errors"
	"gonews/pkg/models"

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

func (db *DB) AddPost(p models.Post) error {

	err := db.pool.QueryRow(context.Background(), `
		INSERT INTO news (title, content, pubtime, link)
		VALUES ($1, $2, $3, $4);
		`,
		p.Title,
		p.Content,
		p.PubTime,
		p.Link,
	).Scan()
	return err
}

func (db *DB) AddPosts(posts []models.Post) error {
	for _, post := range posts {
		err := db.AddPost(post)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) GetPostByID(id int) (models.Post, error) {
	row := db.pool.QueryRow(context.Background(), `	
	SELECT * FROM news 
    WHERE id =$1;
	`, id)

	var post models.Post

	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.PubTime,
		&post.Link)

	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (db *DB) GetPostByHeader(pattern string, limit, offset int) ([]models.Post, models.Pagination, error) {

	pattern = "%" + pattern + "%"

	pagination := models.Pagination{
		Page:  offset/limit + 1,
		Limit: limit,
	}

	row := db.pool.QueryRow(context.Background(), "SELECT count(*) FROM news WHERE title ILIKE $1;", pattern)

	err := row.Scan(&pagination.NumOfPages)

	if pagination.NumOfPages%limit > 0 {
		pagination.NumOfPages = pagination.NumOfPages/limit + 1
	} else {
		pagination.NumOfPages /= limit
	}

	if err != nil {
		return nil, models.Pagination{}, err
	}

	rows, err := db.pool.Query(context.Background(), "SELECT * FROM news WHERE title ILIKE $1 ORDER BY pubtime DESC LIMIT $2 OFFSET $3;", pattern, limit, offset)
	if err != nil {
		return nil, models.Pagination{}, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.PubTime, &p.Link)
		if err != nil {
			return nil, models.Pagination{}, err
		}
		posts = append(posts, p)
	}
	return posts, pagination, rows.Err()
}

func (db *DB) Posts(limit, offset int) ([]models.Post, error) {
	pagination := models.Pagination{
		Page:  offset/limit + 1,
		Limit: limit,
	}
	rows, err := db.pool.Query(context.Background(), `
	SELECT * FROM news
	ORDER BY pubtime DESC LIMIT $1 OFFSET $2
	`,
		pagination.Limit, pagination.Page,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PubTime,
			&p.Link,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}
	return posts, rows.Err()
}
