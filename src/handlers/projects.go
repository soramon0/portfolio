package handlers

import (
	"math"

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
	page := c.QueryInt("page", 1)
	size := c.QueryInt("size", 10)

	if page <= 0 {
		page = 1
	}

	if size > 10 || size <= 0 {
		size = 10
	}

	projects, err := p.store.ListProjectsWithGallery(c.Context(), database.ListPublishedProjectsParams{
		Limit:  int32(size),
		Offset: int32(size) * int32(page-1),
	})

	if err != nil {
		p.log.ErrorF("could not fetch projects: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch projects"}
	}

	count, err := p.store.CountPublishedProjects(c.Context())
	if err != nil {
		p.log.ErrorF("failed to count projects: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch projects"}
	}

	totalPages := int64(math.Ceil(float64(count) / float64(size)))
	return c.JSON(types.NewAPIListResponse(projects, count, totalPages))
}
