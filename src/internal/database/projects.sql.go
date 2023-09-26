// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: projects.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const listProjects = `-- name: ListProjects :many
SELECT
  p.id,
  p.client_name,
  p.name,
  p.description,
  p.live_link,
  p.code_link,
  p.start_date,
  p.end_date,
  p.created_at,
  p.updated_at,
  COALESCE(
    (
      SELECT 
        JSON_AGG(
          JSON_BUILD_OBJECT(
            'id', f.id,
            'url', f.url,
            'alt', f.alt,
            'name', f.name,
            'uploaded_at', f.uploaded_at,
            'type', f.type
          ) 
          ORDER BY f.id
        )
      FROM files AS f WHERE f.project_id = p.id
    )::json,
    '[]'::json
  ) AS gallery
FROM
  projects AS p
ORDER BY
  p.id
`

type ListProjectsRow struct {
	ID          uuid.UUID      `json:"id"`
	ClientName  string         `json:"client_name"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	LiveLink    sql.NullString `json:"live_link"`
	CodeLink    sql.NullString `json:"code_link"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     sql.NullTime   `json:"end_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Gallery     interface{}    `json:"gallery"`
}

func (q *Queries) ListProjects(ctx context.Context) ([]ListProjectsRow, error) {
	rows, err := q.db.QueryContext(ctx, listProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProjectsRow{}
	for rows.Next() {
		var i ListProjectsRow
		if err := rows.Scan(
			&i.ID,
			&i.ClientName,
			&i.Name,
			&i.Description,
			&i.LiveLink,
			&i.CodeLink,
			&i.StartDate,
			&i.EndDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Gallery,
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
