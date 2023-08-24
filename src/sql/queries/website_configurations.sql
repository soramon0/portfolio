
-- name: GetWebsiteConfigurations :many
SELECT * FROM website_configurations ORDER BY configuration_name;

-- name: GetWebsiteConfigurationByName :one
SELECT * FROM website_configurations WHERE configuration_name = $1 LIMIT 1;
