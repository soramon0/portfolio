package handlers

import (
	"github.com/gofiber/fiber/v2"
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
	projects, err := p.store.ListProjectsWithGallery(c.Context())
	if err != nil {
		p.log.ErrorF("could not fetch projects: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch projects"}
	}

	return c.JSON(types.NewAPIListResponse(projects, len(projects)))
}
