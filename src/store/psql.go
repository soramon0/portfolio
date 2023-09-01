package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/goombaio/namegenerator"
	_ "github.com/lib/pq"

	"github.com/soramon0/portfolio/src/internal/database"
)

type Store interface {
	database.Querier
	GenerateUniqueUsername(ctx context.Context, retryCount int) (string, error)
}

type psqlStore struct {
	*database.Queries
}

func NewStore(url string) (Store, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &psqlStore{Queries: database.New(db)}, nil
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
