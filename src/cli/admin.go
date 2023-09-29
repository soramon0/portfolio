package cli

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/store"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

type createAdminPayload struct {
	Email     string `validate:"required,email,omitempty"`
	Password  string `validate:"required,min=6,omitempty"`
	Username  string `validate:"required,min=3,omitempty"`
	Firstname string `validate:"omitempty"`
	Lastname  string `validate:"omitempty"`
}

func newAdminCommand() *cli.Command {
	return &cli.Command{
		Name:  "admin",
		Usage: "admin management",
		Subcommands: []*cli.Command{
			{
				Name:  "create-admin",
				Usage: "database migrations",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "email",
						Aliases:  []string{"e"},
						Usage:    "email for admin user",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "username",
						Aliases:  []string{"u"},
						Usage:    "username for admin user",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "password",
						Aliases:  []string{"p"},
						Usage:    "password for admin user",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "firstname",
						Aliases: []string{"fn"},
						Usage:   "admin user first name",
					},
					&cli.StringFlag{
						Name:    "lastname",
						Aliases: []string{"ln"},
						Usage:   "admin user last name",
					},
				},
				Action: func(ctx *cli.Context) error {
					db, err := store.NewStore(lib.GetDatabaseURL())
					if err != nil {
						return err
					}
					defer db.Close()

					row := db.QueryRow("SELECT EXISTS (SELECT * FROM users WHERE user_type = 'admin' LIMIT 1);")
					var exists bool
					if err := row.Scan(&exists); err != nil {
						return nil
					}
					if exists {
						return errors.New("admin user already exists")
					}

					payload := &createAdminPayload{
						Email:     ctx.String("email"),
						Password:  ctx.String("password"),
						Username:  ctx.String("username"),
						Firstname: ctx.String("firstname"),
						Lastname:  ctx.String("lastname"),
					}

					vt, err := lib.NewValidator()
					if err != nil {
						return err
					}

					if err := vt.Validator.Struct(payload); err != nil {
						var ve validator.ValidationErrors
						if errors.As(err, &ve) {
							return errors.New(vt.ValidationErrors(ve).Errors[0].Message)
						}
						return err
					}

					password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
					if err != nil {
						if err == bcrypt.ErrPasswordTooLong {
							return errors.New("password is too long")
						}
						return err
					}

					createdAt := time.Now().UTC()
					_, err = db.CreateUser(context.Background(), database.CreateUserParams{
						ID:        uuid.New(),
						CreatedAt: createdAt,
						UpdatedAt: createdAt,
						UserType:  database.UserTypeAdmin,
						Email:     strings.Trim(strings.ToLower(payload.Email), " "),
						Password:  string(password),
						Username:  payload.Username,
						FirstName: payload.Firstname,
						LastName:  payload.Lastname,
					})
					return err
				},
			},
		},
	}
}
