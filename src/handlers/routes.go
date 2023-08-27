package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/soramon0/portfolio/src/server"
	"github.com/soramon0/portfolio/src/template"
)

func Register(s *server.AppServer) {
	m := NewMiddleware(s.DB, s.Cache, s.Log)
	m.fiberMiddleware(s.App)

	apiRoutes := s.App.Group("/api").Use(logger.New()).Use(m.WithRateLimit(20, 60, 60))
	apiRoutes.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true})
	})

	v1Router := apiRoutes.Group("/v1")

	authHandlers := NewAuth(s.DB, s.VT, s.Log)
	authRouter := v1Router.Group("/auth")
	authRouter.Post("/register", m.WithWebsiteConfig("allow_user_register", "allow", "registration is disabled"), authHandlers.Register)
	authRouter.Post("/login", m.WithWebsiteConfig("allow_user_login", "allow", "login is disabled"), authHandlers.Login)
	authRouter.Post("/logout", authHandlers.Logout)

	usersHandlers := NewUsers(s.DB, s.Log)
	usersRouter := v1Router.Group("/users").Use(m.WithAuthenticatedUser)
	usersRouter.Get("/me", usersHandlers.GetMe)

	adminUsersRouter := usersRouter.Use(m.WithAuthenticatedAdmin)
	adminUsersRouter.Get("/", usersHandlers.GetUsers)
	adminUsersRouter.Get("/:id", usersHandlers.GetUserById)

	apiRoutes.Use(
		func(c *fiber.Ctx) error {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "Maybe you are lost"}
		},
	)

	// Serve static files
	s.App.All("/*", filesystem.New(filesystem.Config{
		Root:         template.Dist(s.Log),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))
}
