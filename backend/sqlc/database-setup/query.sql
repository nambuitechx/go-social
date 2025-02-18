-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users(id, email) VALUES($1, $2) RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteUserByEmail :exec
DELETE FROM users WHERE email = $1;




-- name: GetPosts :many
SELECT * FROM posts;

-- name: GetPostById :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- name: CreatePost :one
INSERT INTO posts(id, content, user_id) VALUES($1, $2, $3) RETURNING *;

-- name: DeletePostById :exec
DELETE FROM posts WHERE id = $1;

