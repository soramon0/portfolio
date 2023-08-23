package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/internal/types"
	"github.com/soramon0/portfolio/src/lib"
)

type Users struct {
	db  *database.Queries
	log *lib.AppLogger
}

// New Users is used to create a new Users controller.
func NewUsers(db *database.Queries, l *lib.AppLogger) *Users {
	return &Users{
		db:  db,
		log: l,
	}
}

func (u *Users) GetMe(ctx *fiber.Ctx, authUser database.GetUserByIdRow) error {
	return ctx.JSON(authUser)
}

func (u *Users) GetUsers(c *fiber.Ctx, user database.GetUserByIdRow) error {
	users, err := u.db.ListUsers(c.Context())
	if err != nil {
		u.log.ErrorF("could not fetch users: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch users"}
	}

	return c.JSON(types.NewAPIListResponse(users, len(users)))
}

func (u *Users) GetUserById(ctx *fiber.Ctx, authUser database.GetUserByIdRow) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		u.log.Infof("failed to parse id: %v\n", err)
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid user id"}
	}

	user, err := u.db.GetUserById(ctx.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "user not found"}
		}

		u.log.ErrorF("failed to fetch user: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "failed to fetch user"}
	}

	return ctx.JSON(types.NewAPIResponse(user))
}
