package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/soramon0/portfolio/src/internal/database"
)

type Store interface {
	database.Querier

	// used db management cases
	QueryRow(query string, args ...any) pgx.Row
	Exec(query string, args ...any) (pgconn.CommandTag, error)
	Close()
	Migrate(dir string, command string, arguments ...string) error

	// business related methods
	GenerateUniqueUsername(ctx context.Context, retryCount int) (string, error)
	CreateInitialWebsiteConfigs(ctx context.Context, configs []database.CreateWebsiteConfigParams) error
	GetInitialWebsiteConfigParams() []database.CreateWebsiteConfigParams
}

type psqlStore struct {
	*database.Queries
	pool *pgxpool.Pool
}

func NewStore(url string) (Store, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &psqlStore{
		pool:    pool,
		Queries: database.New(pool),
	}, nil
}

func (s *psqlStore) Migrate(dir string, command string, arguments ...string) error {
	db := stdlib.OpenDB(*s.pool.Config().ConnConfig)
	defer db.Close()

	// May need to configure the max idle conns
	// Link: https://github.com/jackc/pgx/blob/163eb68866a76a9cf6a15500303725aac32f6ca3/stdlib/sql.go#L230
	// db.SetMaxIdleConns(0)
	return goose.Run(command, db, dir, arguments...)
}

func (s *psqlStore) QueryRow(query string, args ...any) pgx.Row {
	return s.pool.QueryRow(context.Background(), query, args...)
}

func (s *psqlStore) Exec(query string, args ...any) (pgconn.CommandTag, error) {
	return s.pool.Exec(context.Background(), query, args...)
}

func (s *psqlStore) Close() {
	s.pool.Close()
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
			ID:                 pgtype.UUID{Bytes: uuid.New(), Valid: true},
			Active:             true,
			CreatedAt:          pgtype.Timestamptz{Time: now, Valid: true},
			UpdatedAt:          pgtype.Timestamptz{Time: now, Valid: true},
			Description:        pgtype.Text{String: "", Valid: false},
			ConfigurationName:  "allow_user_login",
			ConfigurationValue: database.WebsiteConfigValueDisallow,
		},
		{
			ID:                 pgtype.UUID{Bytes: uuid.New(), Valid: true},
			Active:             true,
			CreatedAt:          pgtype.Timestamptz{Time: now, Valid: true},
			UpdatedAt:          pgtype.Timestamptz{Time: now, Valid: true},
			Description:        pgtype.Text{String: "", Valid: false},
			ConfigurationName:  "allow_user_register",
			ConfigurationValue: database.WebsiteConfigValueDisallow,
		},
	}
}

func (s *psqlStore) CreateInitialWebsiteConfigs(ctx context.Context, cfgs []database.CreateWebsiteConfigParams) error {
	if len(cfgs) == 0 {
		return fmt.Errorf("website configurations cannot be empty")
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := s.Queries.WithTx(tx)

	for _, c := range cfgs {
		_, err := qtx.GetWebsiteConfigurationByName(ctx, c.ConfigurationName)
		if err != pgx.ErrNoRows {
			return err
		}

		if _, err := qtx.CreateWebsiteConfig(ctx, c); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
