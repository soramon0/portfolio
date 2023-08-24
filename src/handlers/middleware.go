package handlers

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/lib"
)

type Middleware struct {
	db  *database.Queries
	log *lib.AppLogger
}

func NewMiddleware(db *database.Queries, l *lib.AppLogger) *Middleware {
	return &Middleware{
		db:  db,
		log: l,
	}
}

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func (m *Middleware) fiberMiddleware(a *fiber.App) {
	a.Use(
		// Recover from panics
		recover.New(),
		// Add CORS to each route.
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodHead,
				fiber.MethodPut,
				fiber.MethodDelete,
				fiber.MethodPatch,
			}, ","),
			AllowCredentials: true,
		}),
	)
}

const localsUserKey = "user"

func getAuthenticatedUser(ctx *fiber.Ctx) *database.GetUserByIdRow {
	value := ctx.Locals(localsUserKey)
	if user, ok := value.(database.GetUserByIdRow); ok {
		return &user
	}
	return nil
}

func (m *Middleware) WithAuthenticatedUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(lib.GetTokenSecret()), nil
	})
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "unauthenticated"}
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok && !token.Valid {
		m.log.ErrorF("invalid token claims types: %+v\n", token.Claims)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to authenticated"}
	}

	userId, err := uuid.Parse(claims.Issuer)
	if err != nil {
		m.log.ErrorF("failed to parse issuer as uuid: %v\n", err)
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "unauthenticated"}
	}

	user, err := m.db.GetUserById(c.Context(), userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "unauthenticated"}
		}
		m.log.ErrorF("failed to parse issuer as uuid: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to authenticated"}
	}

	c.Locals(localsUserKey, user)

	return c.Next()
}

func (m *Middleware) WithAuthenticatedAdmin(c *fiber.Ctx) error {
	user := getAuthenticatedUser(c)
	if user == nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "unauthentiacted"}
	}
	if user.UserType != "admin" {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "unauthentiacted"}
	}

	return c.Next()
}
