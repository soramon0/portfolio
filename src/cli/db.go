package cli

import (
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/store"
	"github.com/urfave/cli/v2"
)

func newDBCommand() *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database management",
		Subcommands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "database migrations",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "run up migrations",
						Action: func(c *cli.Context) error {
							db, err := store.NewStore(lib.GetDatabaseURL())
							if err != nil {
								return err
							}
							defer db.Close()

							return db.Migrate("./src/sql/schema", "up")
						},
					},
					{
						Name:  "down",
						Usage: "run down migrations",
						Action: func(c *cli.Context) error {
							db, err := store.NewStore(lib.GetDatabaseURL())
							if err != nil {
								return err
							}
							defer db.Close()
							return db.Migrate("./src/sql/schema", "down")
						},
					},
				},
			},
		},
	}
}
