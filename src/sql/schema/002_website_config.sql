-- +goose Up
CREATE TABLE website_configurations (
  id UUID PRIMARY KEY,
  configuration_name VARCHAR(255) UNIQUE NOT NULL,
  configuration_value TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  active BOOLEAN NOT NULL DEFAULT TRUE
);

-- +goose Down
DROP TABLE website_configurations;
