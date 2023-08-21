package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func fiberMiddleware(a *fiber.App) {
	a.Use(
		// Recover from panics
		recover.New(),
		// Add CORS to each route.
		cors.New(),
	)
}
