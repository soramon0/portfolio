-- +goose Up

-- +goose StatementBegin
DO
$$
  BEGIN
    CREATE TYPE USER_TYPE AS ENUM ('admin', 'user');
  EXCEPTION
    WHEN duplicate_object THEN null;
  END
$$;
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  user_type USER_TYPE NOT NULL DEFAULT 'user',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users ADD CONSTRAINT user_type_check CHECK (user_type IN ('admin', 'user'));

-- +goose Down
ALTER TABLE users DROP CONSTRAINT user_type_check;
DROP TABLE IF EXISTS users;
