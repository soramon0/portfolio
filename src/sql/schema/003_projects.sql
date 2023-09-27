-- +goose Up
CREATE TABLE IF NOT EXISTS projects (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  client_name VARCHAR(255) NOT NULL,
  live_link TEXT,
  code_link TEXT,
  start_date DATE NOT NULL,
  end_date Date,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE FILE_TYPE AS ENUM ('image', 'document');

CREATE TABLE IF NOT EXISTS files (
  id UUID PRIMARY KEY,
  url TEXT NOT NULL,
  alt TEXT NOT NULL,
  name TEXT,
  type FILE_TYPE NOT NULL DEFAULT 'image',
  uploaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Project(files): Many to one
-- projects can have zero or many images for gallery
ALTER TABLE files
ADD COLUMN project_id UUID,
ADD CONSTRAINT fk_files_projects FOREIGN KEY (project_id)
REFERENCES projects(id) ON DELETE CASCADE;

-- Project(file): One to one
-- projects can have one cover image
ALTER TABLE projects
ADD COLUMN cover_image_id UUID UNIQUE,
ADD CONSTRAINT fk_project_file FOREIGN KEY (cover_image_id)
REFERENCES files(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE files
DROP CONSTRAINT IF EXISTS fk_files_projects;

ALTER TABLE projects
DROP CONSTRAINT IF EXISTS fk_project_file;

DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS files;
DROP TYPE FILE_TYPE;
