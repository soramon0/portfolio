package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/soramon0/portfolio/src/internal/database"
)

type ProjectWithGallary struct {
	database.GetPublishedProjectBySlugRow
	Gallery []database.File `json:"gallery"`
}

func (s *psqlStore) GetPublishedProject(ctx context.Context, slug string) (ProjectWithGallary, error) {
	var project ProjectWithGallary

	row, err := s.Queries.GetPublishedProjectBySlug(ctx, slug)
	if err != nil {
		return project, err
	}

	v, ok := row.Gallery.([]byte)
	if !ok {
		return project, fmt.Errorf("failed to cast gallery(%T) to bytes", row.Gallery)
	}

	if err := json.Unmarshal(v, &project.Gallery); err != nil {
		return project, err
	}

	project.ID = row.ID
	project.ClientName = row.ClientName
	project.Name = row.Name
	project.Slug = row.Slug
	project.Subtitle = row.Subtitle
	project.Description = row.Description
	project.LiveLink = row.LiveLink
	project.CodeLink = row.CodeLink
	project.StartDate = row.StartDate
	project.EndDate = row.EndDate
	project.LaunchDate = row.LaunchDate
	project.CreatedAt = row.CreatedAt
	project.UpdatedAt = row.UpdatedAt
	project.CoverImageName = row.CoverImageName
	project.CoverImageUrl = row.CoverImageUrl
	project.CoverImageAlt = row.CoverImageAlt

	return project, nil
}
