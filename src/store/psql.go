package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/goombaio/namegenerator"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/soramon0/portfolio/src/internal/database"
)

type Store interface {
	database.Querier
	GenerateUniqueUsername(ctx context.Context, retryCount int) (string, error)
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	Close() error
	Migrate(dir string, command string, arguments ...string) error
}

type psqlStore struct {
	*database.Queries
	db *sql.DB
}

func NewStore(url string) (Store, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &psqlStore{
		db:      db,
		Queries: database.New(db),
	}, nil
}

func (s *psqlStore) Migrate(dir string, command string, arguments ...string) error {
	return goose.Run(command, s.db, dir, arguments...)
}

func (s *psqlStore) QueryRow(query string, args ...any) *sql.Row {
	return s.db.QueryRow(query, args...)
}

func (s *psqlStore) Exec(query string, args ...any) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

func (s *psqlStore) Close() error {
	return s.db.Close()
}

func (s *psqlStore) GenerateUniqueUsername(ctx context.Context, retryCount int) (string, error) {
	seed := time.Now().UTC().UnixNano()
	username := namegenerator.NewNameGenerator(seed).Generate()
	exists, err := s.CheckUserExistsByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if !exists {
		return username, nil
	}
	if retryCount <= 0 {
		return "", errors.New("failed to generate unique username")
	}
	return s.GenerateUniqueUsername(ctx, retryCount-1)
}
