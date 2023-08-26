package main

import (
	"github.com/soramon0/portfolio/src/cache"
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
	cache, err := cache.NewCache(lib.GetRedisURL())
	if err != nil {
		logger.ErrorFatalF("could not connect to redis: %v", err)
	}

	vt, err := lib.NewValidator()
	if err != nil {
		logger.ErrorFatalF("could not create validator: %v", err)
	}

	appServer := lib.NewAppServer(app, db, cache, vt, logger)

	handlers.Register(appServer)
	appServer.StartServer()
}
