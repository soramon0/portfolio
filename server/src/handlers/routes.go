package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/lib"
)

func Register(a *fiber.App, db *database.Queries, l *lib.AppLogger) {
	fiberMiddleware(a)

	a.Get("/api/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true})
	})

	usersHandlers := NewUsers(db, l)
	usersRouter := a.Group("/api/v1/users")
	usersRouter.Get("/", usersHandlers.GetUsers)
	usersRouter.Get("/:id", usersHandlers.GetUserById)
}
