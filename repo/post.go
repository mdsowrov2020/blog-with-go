package repo

import (
	"database/sql"

	"blog/domain"
	"blog/post"

	"github.com/jmoiron/sqlx"
)

type PostRepo interface {
	post.PostRepo
}

type postRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) PostRepo {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Create(p domain.Post) (*domain.Post, error) {
	query := `
	INSERT INTO posts (
	 title,
	 description,
	 image_url
	) VALUES (
	 $1,
	 $2,
	 $3
	 ) RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.ImageURL)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *postRepo) List() ([]*domain.Post, error) {
	postList := []*domain.Post{}

	query := `
	SELECT id, title, description, image_url FROM posts
	`

	err := r.db.Select(&postList, query)
	if err != nil {
		return nil, err
	}

	return postList, nil
}

func (r *postRepo) Get(id int) (*domain.Post, error) {
	var post domain.Post

	query := `
	SELECT 
	id,
	title,
	description,
	image_url
	FROM posts
	WHERE id = $1
	`
	err := r.db.Get(&post, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &post, nil
}

func (r *postRepo) Update(p domain.Post) (*domain.Post, error) {
	query := `
	UPDATE posts
	SET title = $1, description = $2,image_url = $3
	WHERE id = $4

	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.ImageURL, p.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *postRepo) Delete(id int) error {
	query := `
	DELETE FROM posts
	WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
