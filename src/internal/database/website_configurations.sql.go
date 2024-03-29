// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: website_configurations.sql

package database

import (
	"context"
)

const getWebsiteConfigurationByName = `-- name: GetWebsiteConfigurationByName :one
SELECT configuration_name, configuration_value, description, created_at, updated_at, active FROM website_configurations WHERE configuration_name = $1 LIMIT 1
`

func (q *Queries) GetWebsiteConfigurationByName(ctx context.Context, configurationName string) (WebsiteConfiguration, error) {
	row := q.db.QueryRowContext(ctx, getWebsiteConfigurationByName, configurationName)
	var i WebsiteConfiguration
	err := row.Scan(
		&i.ConfigurationName,
		&i.ConfigurationValue,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
	)
	return i, err
}

const getWebsiteConfigurations = `-- name: GetWebsiteConfigurations :many
SELECT configuration_name, configuration_value, description, created_at, updated_at, active FROM website_configurations ORDER BY configuration_name
`

func (q *Queries) GetWebsiteConfigurations(ctx context.Context) ([]WebsiteConfiguration, error) {
	rows, err := q.db.QueryContext(ctx, getWebsiteConfigurations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []WebsiteConfiguration{}
	for rows.Next() {
		var i WebsiteConfiguration
		if err := rows.Scan(
			&i.ConfigurationName,
			&i.ConfigurationValue,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
