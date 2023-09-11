
-- name: GetWebsiteConfigurations :many
SELECT * FROM website_configurations ORDER BY configuration_name;

-- name: GetWebsiteConfigurationByName :one
SELECT * FROM website_configurations WHERE configuration_name = $1 LIMIT 1;

-- name: CreateWebsiteConfig :one
INSERT INTO website_configurations (id, created_at, updated_at, configuration_name, configuration_value, description, active)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
