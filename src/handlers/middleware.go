package handlers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/soramon0/portfolio/src/cache"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/lib"
)

type Middleware struct {
	db    *database.Queries
	cache *cache.Cache
	log   *lib.AppLogger
}

func NewMiddleware(db *database.Queries, cache *cache.Cache, l *lib.AppLogger) *Middleware {
	return &Middleware{
		db:    db,
		cache: cache,
		log:   l,
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
	if !ok || !token.Valid {
		m.log.ErrorF("invalid token: %+v\n", token)
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

func (m *Middleware) WithWebsiteConfig(name string, value string, errMsg string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		wc, err := m.db.GetWebsiteConfigurationByName(c.Context(), name)
		if err != nil {
			if err == sql.ErrNoRows {
				return &fiber.Error{Code: fiber.StatusUnauthorized, Message: errMsg}
			}
			m.log.ErrorF("failed to get website config for %s: %v\n", name, err)
			return &fiber.Error{Code: fiber.StatusInternalServerError, Message: errMsg}
		}

		if !wc.Active || strings.ToLower(wc.ConfigurationValue) != strings.ToLower(value) {
			return &fiber.Error{Code: fiber.StatusUnauthorized, Message: errMsg}
		}

		return c.Next()
	}
}

func (m *Middleware) WithRateLimit(limit int, perSec int, backOffDuration int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := "ratelimit:user:" + c.IP()
		accept, duration := m.cache.CounterRateLimit(c.Context(), key, limit, perSec, backOffDuration)
		if !accept {
			msg := fmt.Sprintf("Too many requests. Please try again after %ds", duration)
			return &fiber.Error{Code: fiber.StatusTooManyRequests, Message: msg}
		}
		return c.Next()
	}
}
