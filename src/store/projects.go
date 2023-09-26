package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/soramon0/portfolio/src/internal/database"
)

const listProjects = `
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
        SELECT JSON_AGG(
          JSON_BUILD_OBJECT(
            'id', f.id,
            'url', f.url,
            'alt', f.alt,
            'name', f.name,
            'uploaded_at', f.uploaded_at,
            'type', f.type
          ) ORDER BY f.id
        )
        FROM files AS f
        WHERE f.project_id = p.id
      )::json,
      '[]'::json
    ) AS gallery
FROM
    projects AS p
ORDER BY
    p.id;
`

type ProjectWithGallary struct {
	ID          uuid.UUID       `json:"id"`
	ClientName  string          `json:"client_name"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	LiveLink    sql.NullString  `json:"live_link"`
	CodeLink    sql.NullString  `json:"code_link"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     sql.NullTime    `json:"end_date"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Gallery     []database.File `json:"gallery"`
}

func (s *psqlStore) ListProjectsWithGallery(ctx context.Context) ([]ProjectWithGallary, error) {
	rows, err := s.db.QueryContext(ctx, listProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProjectWithGallary{}
	for rows.Next() {
		var i ProjectWithGallary
		var imagesJson []byte
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
			&imagesJson,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(imagesJson, &i.Gallery); err != nil {
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
