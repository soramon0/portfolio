package main

import (
	"os"

	"github.com/soramon0/portfolio/cli"
	"github.com/soramon0/portfolio/src/cache"
	"github.com/soramon0/portfolio/src/configs"
	"github.com/soramon0/portfolio/src/handlers"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/server"
	"github.com/soramon0/portfolio/src/store"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if len(os.Args) >= 2 {
		cli.NewCli(lib.NewLogger()).Run(os.Args)
	}

	logger := lib.NewLogger()
	app := fiber.New(configs.FiberConfig())
	db, err := store.NewStore(lib.GetDatabaseURL())
	if err != nil {
		logger.ErrorFatalF("could not connect to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.ErrorFatalF("failed to close db connection: %v", err)
		}
	}()
	if err := db.Migrate("./src/sql/schema", "up"); err != nil {
		logger.ErrorFatalF("failed to migrate db: %v", err)
	}

	cache, err := cache.NewCache(lib.GetRedisURL(), logger)
	if err != nil {
		logger.ErrorFatalF("could not connect to redis: %v", err)
	}

	vt, err := lib.NewValidator()
	if err != nil {
		logger.ErrorFatalF("could not create validator: %v", err)
	}

	appServer := server.NewAppServer(app, db, cache, vt, logger)

	handlers.Register(appServer)
	appServer.StartServer()
}
