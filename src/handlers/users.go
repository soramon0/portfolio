package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/soramon0/portfolio/src/internal/types"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/store"
)

type Users struct {
	store store.Store
	log   *lib.AppLogger
}

// New Users is used to create a new Users controller.
func NewUsers(s store.Store, l *lib.AppLogger) *Users {
	return &Users{
		store: s,
		log:   l,
	}
}

func (u *Users) GetMe(ctx *fiber.Ctx) error {
	user := getAuthenticatedUser(ctx)
	if user == nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "unauthentiacted"}
	}
	return ctx.JSON(types.NewAPIResponse(user))
}

func (u *Users) GetUsers(c *fiber.Ctx) error {
	users, err := u.store.ListUsers(c.Context())
	if err != nil {
		u.log.ErrorF("could not fetch users: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch users"}
	}

	return c.JSON(types.NewAPIListResponse(users, len(users)))
}

func (u *Users) GetUserById(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		u.log.Infof("failed to parse id: %v\n", err)
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid user id"}
	}

	user, err := u.store.GetUserById(ctx.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "user not found"}
		}

		u.log.ErrorF("failed to fetch user: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch user"}
	}

	return ctx.JSON(types.NewAPIResponse(user))
}
