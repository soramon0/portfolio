package main

import (
	"github.com/soramon0/portfolio/src/configs"
	"github.com/soramon0/portfolio/src/handlers"
	"github.com/soramon0/portfolio/src/lib"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logger := lib.NewLogger()
	db := lib.NewDB(lib.GetDatabaseURL(), logger)
	app := fiber.New(configs.FiberConfig())

	vt, err := lib.NewValidator()
	if err != nil {
		logger.ErrorFatalF("could not create validator %v", err)
	}

	handlers.Register(app, db, vt, logger)
	lib.StartServer(app, logger)
}
