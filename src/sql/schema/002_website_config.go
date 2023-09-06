package schema

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up002WebsiteConfig, Down002WebsiteConfig)
}

func Up002WebsiteConfig(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE website_configurations (
  configuration_name VARCHAR(255) UNIQUE NOT NULL,
  configuration_value TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  active BOOLEAN NOT NULL DEFAULT TRUE
);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func Down002WebsiteConfig(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE website_configurations;")
	return err
}
