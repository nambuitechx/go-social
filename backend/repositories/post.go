package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"

	"github.com/nambuitechx/go-social/models"
)

type PostRepository struct {
	DB *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{ DB: db }
}

func (r *PostRepository) SelectPosts(limit *int, offset *int) []models.PostModel {
	posts := []models.PostModel{}

	if *limit < 0 {
		statement := "SELECT id, content, user_id, created_at, updated_at FROM posts"
		r.DB.Select(&posts, statement)
	} else {
		statement := "SELECT id, content, user_id, created_at, updated_at FROM posts LIMIT $1 OFFSET $2"
		r.DB.Select(&posts, statement, *limit, *offset)
	}
	
	return posts
}

func (r *PostRepository) SelectPostById(id *string) (*models.PostModel, error) {
	post := &models.PostModel{}

	statement := "SELECT id, content, user_id, created_at, updated_at FROM posts WHERE id = $1"
	err := r.DB.Get(post, statement, *id)
	return post, err
}

func (r *PostRepository) InsertPost(payload *models.CreatePostPayload, userId string) (*models.PostModel, error) {
	var post = models.PostModel{}
	statement := "INSERT INTO posts(id, content, user_id) VALUES($1, $2, $3) RETURNING id, content, user_id, created_at, updated_at"
	err := r.DB.Get(&post, statement, uuid.New().String(), payload.Content, userId)
	return &post, err
}

func (r *PostRepository) DeletePostById(id *string) error {
	statement := "DELETE FROM posts WHERE id = $1"
	_, err := r.DB.Exec(statement, *id)
	return err
}
