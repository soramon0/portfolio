package schema

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up001Users, Down001Users)
}

func Up001Users(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE users (
  id UUID PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  user_type VARCHAR(255) NOT NULL DEFAULT 'user',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}

	query = "ALTER TABLE users ADD CONSTRAINT user_type_check CHECK (user_type IN ('admin', 'user'));"
	_, err := tx.ExecContext(ctx, query)
	return err
}

func Down001Users(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE users;")
	return err
}
