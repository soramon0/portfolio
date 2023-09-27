package store

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/soramon0/portfolio/src/internal/database"
)

type ProjectWithGallary struct {
	database.ListProjectsRow
	Gallery []database.File `json:"gallery"`
}

func (s *psqlStore) ListProjectsWithGallery(ctx context.Context) ([]ProjectWithGallary, error) {
	rows, err := s.Queries.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	projects := make([]ProjectWithGallary, 0, len(rows))
	for _, row := range rows {
		var i ProjectWithGallary

		i.ID = row.ID
		i.ClientName = row.ClientName
		i.Name = row.Name
		i.Description = row.Description
		i.LiveLink = row.LiveLink
		i.CodeLink = row.CodeLink
		i.StartDate = row.StartDate
		i.EndDate = row.EndDate
		i.CreatedAt = row.CreatedAt
		i.UpdatedAt = row.UpdatedAt
		i.CoverImageName = row.CoverImageName
		i.CoverImageUrl = row.CoverImageUrl
		i.CoverImageAlt = row.CoverImageAlt

		v, ok := row.Gallery.([]byte)
		if !ok {
			return nil, errors.New("failed to convert gallery to bytes")
		}

		if err := json.Unmarshal(v, &i.Gallery); err != nil {
			return nil, err
		}

		projects = append(projects, i)
	}

	return projects, nil
}
