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
		return project, fmt.Errorf("failed to convert gallery(%T) to bytes", row.Gallery)
	}

	if err := json.Unmarshal(v, &project.Gallery); err != nil {
		return project, err
	}

	return project, nil
}
