package store

import (
	"context"
	"encoding/json"

	"github.com/soramon0/portfolio/src/internal/database"
)

type ProjectWithGallary struct {
	database.Project
	Gallery []database.File `json:"gallery"`
}

func (s *psqlStore) ListProjectsWithGallery(ctx context.Context) ([]ProjectWithGallary, error) {
	rows, err := s.db.QueryContext(ctx, database.ListProjects)
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
