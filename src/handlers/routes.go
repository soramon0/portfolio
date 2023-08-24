package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/template"
)

func Register(a *fiber.App, db *database.Queries, vt *lib.ValidatorTranslator, l *lib.AppLogger) {
	middleware := NewMiddleware(db, l)
	middleware.fiberMiddleware(a)

	apiRoutes := a.Group("/api").Use(logger.New())
	apiRoutes.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true})
	})

	v1Router := apiRoutes.Group("/v1")

	authHandlers := NewAuth(db, vt, l)
	authRouter := v1Router.Group("/auth")
	authRouter.Post("/register", authHandlers.Register)
	authRouter.Post("/login", authHandlers.Login)
	authRouter.Post("/logout", authHandlers.Logout)

	usersHandlers := NewUsers(db, l)
	usersRouter := v1Router.Group("/users").Use(middleware.WithAuthenticatedUser)
	usersRouter.Get("/me", usersHandlers.GetMe)

	adminUsersRouter := usersRouter.Use(middleware.WithAuthenticatedAdmin)
	adminUsersRouter.Get("/", usersHandlers.GetUsers)
	adminUsersRouter.Get("/:id", usersHandlers.GetUserById)

	apiRoutes.Use(
		func(c *fiber.Ctx) error {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "Maybe you are lost"}
		},
	)

	// Serve static files
	a.All("/*", filesystem.New(filesystem.Config{
		Root:         template.Dist(l),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))
}
