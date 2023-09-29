package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/soramon0/portfolio/src/internal/database"
)

type ProjectWithGallary struct {
	database.ListPublishedProjectsRow
	Gallery []database.File `json:"gallery"`
}

func (s *psqlStore) ListProjectsWithGallery(ctx context.Context, arg database.ListPublishedProjectsParams) ([]ProjectWithGallary, error) {
	rows, err := s.Queries.ListPublishedProjects(ctx, arg)
	if err != nil {
		return nil, err
	}

	projects := make([]ProjectWithGallary, len(rows))
	for i, row := range rows {
		projects[i].ID = row.ID
		projects[i].ClientName = row.ClientName
		projects[i].Name = row.Name
		projects[i].Slug = row.Slug
		projects[i].Subtitle = row.Subtitle
		projects[i].Description = row.Description
		projects[i].LiveLink = row.LiveLink
		projects[i].CodeLink = row.CodeLink
		projects[i].StartDate = row.StartDate
		projects[i].EndDate = row.EndDate
		projects[i].LaunchDate = row.LaunchDate
		projects[i].CreatedAt = row.CreatedAt
		projects[i].UpdatedAt = row.UpdatedAt
		projects[i].CoverImageName = row.CoverImageName
		projects[i].CoverImageUrl = row.CoverImageUrl
		projects[i].CoverImageAlt = row.CoverImageAlt

		v, ok := row.Gallery.([]byte)
		if !ok {
			return nil, fmt.Errorf("failed to convert gallery(%T) to bytes", row.Gallery)
		}

		if err := json.Unmarshal(v, &projects[i].Gallery); err != nil {
			return nil, err
		}
	}

	return projects, nil
}
