package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/portfolio/src/handlers/paginator"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/internal/types"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/store"
)

type Categories struct {
	store store.Store
	log   *lib.AppLogger
}

// NewCategories is used to create a new categories controller.
func NewCategories(s store.Store, l *lib.AppLogger) *Categories {
	return &Categories{
		store: s,
		log:   l,
	}
}

func (ca *Categories) GetCategories(c *fiber.Ctx) error {
	paginatorHeader := c.Get("X-Paginator", string(paginator.OffsetPaginatorType))
	paginator := paginator.NewPaginator[[]database.Category](
		paginator.ParsePaginatorType(paginatorHeader),
		c.QueryInt("page", 1),
		c.QueryInt("size", 10),
	)

	result, err := paginator.Paginate(func(limit, offset int) ([]database.Category, int64, error) {
		projects, err := ca.store.ListCategories(c.Context(), database.ListCategoriesParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})

		if err != nil {
			ca.log.ErrorF("could not fetch categories: %v\n", err)
			return nil, 0, errors.New("failed to fetch categories")
		}

		count, err := ca.store.CountPublishedProjects(c.Context())
		if err != nil {
			ca.log.ErrorF("failed to count categories: %v\n", err)
			return nil, 0, errors.New("failed to fetch categories")
		}

		return projects, count, nil
	})

	if err != nil {
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return c.JSON(types.NewAPIListResponse(result.Data, result.Count, result.TotalPages))
}
