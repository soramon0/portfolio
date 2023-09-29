package cli

import (
	"os"

	"github.com/soramon0/portfolio/src/lib"
	"github.com/urfave/cli/v2"
)

type Cli struct {
	app *cli.App
	log *lib.AppLogger
}

func NewCli(l *lib.AppLogger) *Cli {
	app := &cli.App{
		Name:        "portfolio",
		Description: "CLI for managing soramon0 portfolio",
		Commands: []*cli.Command{
			newDBCommand(),
			newAdminCommand(),
		},
	}

	return &Cli{app: app, log: l}
}

func (c *Cli) Run(args []string) {
	if err := c.app.Run(os.Args); err != nil {
		c.log.ErrorFatal(err)
	}
	os.Exit(0)
}
