// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package sqlc

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts(id, content, user_id) VALUES($1, $2, $3) RETURNING id, content, user_id, created_at, updated_at
`

type CreatePostParams struct {
	ID      string
	Content string
	UserID  string
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost, arg.ID, arg.Content, arg.UserID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users(id, email) VALUES($1, $2) RETURNING id, email, created_at, updated_at
`

type CreateUserParams struct {
	ID    string
	Email string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePostById = `-- name: DeletePostById :exec
DELETE FROM posts WHERE id = $1
`

func (q *Queries) DeletePostById(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deletePostById, id)
	return err
}

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec
DELETE FROM users WHERE email = $1
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, deleteUserByEmail, email)
	return err
}

const deleteUserById = `-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUserById(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteUserById, id)
	return err
}

const getPostById = `-- name: GetPostById :one
SELECT id, content, user_id, created_at, updated_at FROM posts WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPostById(ctx context.Context, id string) (Post, error) {
	row := q.db.QueryRow(ctx, getPostById, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, content, user_id, created_at, updated_at FROM posts
`

func (q *Queries) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.Query(ctx, getPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, email, created_at, updated_at FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
