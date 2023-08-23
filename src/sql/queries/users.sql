-- name: ListUsers :many
SELECT * FROM users ORDER BY username;

-- name: GetUserById :one
SELECT id, username, email, created_at, updated_at, user_type FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CheckUserExistsByEmail :one
SELECT EXISTS (
  SELECT 1 FROM users WHERE email = $1
);

-- name: CheckUserExistsByUsername :one
SELECT EXISTS (
  SELECT 1 FROM users WHERE username = $1
);

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, email, password, first_name, last_name, user_type)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

