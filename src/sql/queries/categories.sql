-- name: ListCategories :many
SELECT * FROM categories
LIMIT $1 OFFSET $2;

-- name: CountCategories :one
SELECT count(*) FROM categories; 
