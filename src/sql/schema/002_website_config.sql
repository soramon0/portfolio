-- +goose Up
CREATE TYPE WEBSITE_CONFIG_VALUE AS ENUM ('allow', 'disallow');

CREATE TABLE website_configurations (
  id UUID PRIMARY KEY,
  configuration_name VARCHAR(255) UNIQUE NOT NULL,
  configuration_value WEBSITE_CONFIG_VALUE NOT NULL DEFAULT 'disallow',
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  active BOOLEAN NOT NULL DEFAULT TRUE
);

-- +goose Down
DROP TABLE website_configurations;
