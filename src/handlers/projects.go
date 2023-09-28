package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/internal/types"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/store"
)

type Projects struct {
	store store.Store
	log   *lib.AppLogger
}

// New Users is used to create a new Users controller.
func NewProjects(s store.Store, l *lib.AppLogger) *Projects {
	return &Projects{
		store: s,
		log:   l,
	}
}

func (p *Projects) GetProjects(c *fiber.Ctx) error {
	paginator := NewOffsetPaginator[[]store.ProjectWithGallary](c.QueryInt("page", 1), c.QueryInt("size", 10))

	result, err := paginator.Paginate(func(limit, offset int) ([]store.ProjectWithGallary, int64, error) {
		projects, err := p.store.ListProjectsWithGallery(c.Context(), database.ListPublishedProjectsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})

		if err != nil {
			p.log.ErrorF("could not fetch projects: %v\n", err)
			return nil, 0, &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch projects"}
		}

		count, err := p.store.CountPublishedProjects(c.Context())
		if err != nil {
			p.log.ErrorF("failed to count projects: %v\n", err)
			return nil, 0, &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch projects"}
		}

		return projects, count, nil
	})

	if err != nil {
		return err
	}

	return c.JSON(types.NewAPIListResponse(result.Data, result.Count, result.TotalPages))
}
