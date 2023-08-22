// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const checkUserExistsByEmail = `-- name: CheckUserExistsByEmail :one
SELECT EXISTS (
  SELECT 1 FROM users WHERE email = $1
)
`

func (q *Queries) CheckUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUserExistsByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkUserExistsByUsername = `-- name: CheckUserExistsByUsername :one
SELECT EXISTS (
  SELECT 1 FROM users WHERE username = $1
)
`

func (q *Queries) CheckUserExistsByUsername(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUserExistsByUsername, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, email, password, first_name, last_name, user_type)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, username, email, password, first_name, last_name, user_type, created_at, updated_at
`

type CreateUserParams struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserType  string    `json:"user_type"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.UserType,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.UserType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, first_name, last_name, user_type, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.UserType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, email, password, first_name, last_name, user_type, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.UserType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, password, first_name, last_name, user_type, created_at, updated_at FROM users ORDER BY username
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.FirstName,
			&i.LastName,
			&i.UserType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
