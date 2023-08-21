-- name: ListUsers :many
SELECT * FROM users ORDER BY name;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
