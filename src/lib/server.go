package lib

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/portfolio/src/cache"
	"github.com/soramon0/portfolio/src/internal/database"
)

type AppServer struct {
	App   *fiber.App
	DB    *database.Queries
	Cache *cache.Cache
	VT    *ValidatorTranslator
	Log   *AppLogger
}

func NewAppServer(app *fiber.App, db *database.Queries, r *cache.Cache, vt *ValidatorTranslator, l *AppLogger) *AppServer {
	return &AppServer{
		App:   app,
		DB:    db,
		Cache: r,
		VT:    vt,
		Log:   l,
	}
}

func (a *AppServer) StartServer() {
	if GetDevelopmentMode() == "development" {
		listenAndServe(a.App, a.Log)
	} else {
		go listenAndServe(a.App, a.Log)
		startServerWithGracefulShutdown(a.App, a.Log)
	}
}

func listenAndServe(a *fiber.App, l *AppLogger) {
	if err := a.Listen(GetServerBindAddress()); err != nil {
		l.ErrorFatalF("Oops... Server is not running! Reason: %v\n", err)
	}
}

func startServerWithGracefulShutdown(a *fiber.App, l *AppLogger) {
	// trap interupt, sigterm or sighub and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	// Block until a signal is received.
	sig := <-sigChan
	l.Infof("Recieved %s, graceful shutdown...\n", sig)

	// gracefully shutdown the server
	if err := a.Shutdown(); err != nil {
		l.ErrorFatal(err)
	}
}
