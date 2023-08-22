-- +goose Up
CREATE TABLE users (
  id UUID PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  user_type VARCHAR(255) NOT NULL DEFAULT 'user',
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

ALTER TABLE users ADD CONSTRAINT user_type_check CHECK (user_type IN ('admin', 'user'));

-- +goose Down
DROP TABLE users;
