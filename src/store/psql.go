package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/soramon0/portfolio/src/internal/database"
)

type Store interface {
	database.Querier
	GenerateUniqueUsername(ctx context.Context, retryCount int) (string, error)
	CreateInitialWebsiteConfigs(ctx context.Context, configs []database.CreateWebsiteConfigParams) error
	GetInitialWebsiteConfigParams() []database.CreateWebsiteConfigParams
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

func (s *psqlStore) GetInitialWebsiteConfigParams() []database.CreateWebsiteConfigParams {
	now := time.Now()
	return []database.CreateWebsiteConfigParams{
		{
			ID:                 uuid.New(),
			Active:             true,
			CreatedAt:          now,
			UpdatedAt:          now,
			Description:        sql.NullString{},
			ConfigurationName:  "allow_user_login",
			ConfigurationValue: "disallow",
		},
		{
			ID:                 uuid.New(),
			Active:             true,
			CreatedAt:          now,
			UpdatedAt:          now,
			Description:        sql.NullString{},
			ConfigurationName:  "allow_user_register",
			ConfigurationValue: "disallow",
		},
	}
}

func (s *psqlStore) CreateInitialWebsiteConfigs(ctx context.Context, cfgs []database.CreateWebsiteConfigParams) error {
	if len(cfgs) == 0 {
		return fmt.Errorf("website configurations cannot be empty")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := s.Queries.WithTx(tx)

	for _, c := range cfgs {
		_, err := qtx.GetWebsiteConfigurationByName(ctx, c.ConfigurationName)
		if err != sql.ErrNoRows {
			return err
		}

		if _, err := qtx.CreateWebsiteConfig(ctx, c); err != nil {
			return err
		}
	}

	return tx.Commit()
}
