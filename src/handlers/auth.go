package handlers

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/internal/types"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/store"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	store store.Store
	vt    *lib.ValidatorTranslator
	log   *lib.AppLogger
}

func NewAuth(s store.Store, vt *lib.ValidatorTranslator, l *lib.AppLogger) *Auth {
	return &Auth{
		store: s,
		vt:    vt,
		log:   l,
	}
}

type authPayload struct {
	Email    string `json:"email" validate:"required,email,omitempty"`
	Password string `json:"password" validate:"required,omitempty"`
}

func (a *Auth) Register(c *fiber.Ctx) error {
	payload := new(authPayload)
	if err := c.BodyParser(payload); err != nil {
		a.log.ErrorF("failed body valdiation %v\n", err)
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid credentials"}
	}

	if err := a.vt.Validator.Struct(payload); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return c.Status(fiber.StatusBadRequest).JSON(a.vt.ValidationErrors(ve))
		}
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	email := strings.Trim(strings.ToLower(payload.Email), " ")
	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	if err != nil {
		if err == bcrypt.ErrPasswordTooLong {
			return c.Status(fiber.StatusBadRequest).JSON(types.APIValidationErrors{
				Errors: []types.APIFieldError{
					{
						Field:   "password",
						Message: "password is too long",
					},
				},
			})
		}
		a.log.ErrorF("failed to hash password: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "registration failed"}
	}

	userExists, err := a.store.CheckUserExistsByEmail(c.Context(), email)
	if err != nil {
		a.log.ErrorF("failed get user by email: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "registration failed"}
	}

	if userExists {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid credentials"}
	}

	username, err := a.store.GenerateUniqueUsername(c.Context(), 5)
	if err != nil {
		a.log.ErrorF("failed generating unique username: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "registration failed"}
	}
	createdAt := time.Now().UTC()
	user, err := a.store.CreateUser(c.Context(), database.CreateUserParams{
		ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: createdAt, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: createdAt, Valid: true},
		Email:     email,
		Password:  string(password),
		Username:  username,
		UserType:  database.UserTypeUser,
		FirstName: "",
		LastName:  "",
	})
	if err != nil {
		a.log.ErrorF("failed creating user: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "registration failed"}
	}

	return c.JSON(types.NewAPIResponse(user))
}

func (a *Auth) Login(c *fiber.Ctx) error {
	payload := new(authPayload)
	if err := c.BodyParser(payload); err != nil {
		a.log.ErrorF("failed body valdiation %v\n", err)
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid credentials"}
	}

	if err := a.vt.Validator.Struct(payload); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return c.Status(fiber.StatusBadRequest).JSON(a.vt.ValidationErrors(ve))
		}
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()}
	}

	email := strings.Trim(strings.ToLower(payload.Email), " ")

	user, err := a.store.GetUserByEmail(c.Context(), email)
	if err != nil {
		if err == sql.ErrNoRows {
			return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid credentials"}
		}
		a.log.ErrorF("failed get user by email: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "login failed"}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		a.log.Infof("failed password check: %v\n", err)
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "invalid credentials"}
	}

	id, err := uuid.FromBytes(user.ID.Bytes[:])
	if err != nil {
		a.log.Infof("failed to parse user id: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "login failed"}
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		Issuer:    id.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(lib.GetTokenSecret()))
	if err != nil {
		a.log.ErrorF("failed to sign token: %v\n", err)
		return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "login failed"}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    ss,
		Expires:  expiresAt,
		HTTPOnly: true,
	})

	return c.JSON(types.NewAPIResponse(user))
}

func (a *Auth) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(types.NewAPIResponse(fiber.Map{"logout": true}))
}
