package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/soramon0/portfolio/src/cache"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/store"
)

type AppServer struct {
	App   *fiber.App
	Store store.Store
	Cache *cache.Cache
	VT    *lib.ValidatorTranslator
	Log   *lib.AppLogger
}

func NewAppServer(app *fiber.App, s store.Store, r *cache.Cache, vt *lib.ValidatorTranslator, l *lib.AppLogger) *AppServer {
	return &AppServer{
		App:   app,
		Store: s,
		Cache: r,
		VT:    vt,
		Log:   l,
	}
}

func (a *AppServer) StartServer() {
	wConfigs := a.Store.GetInitialWebsiteConfigParams()
	if err := a.Store.CreateInitialWebsiteConfigs(context.Background(), wConfigs); err != nil {
		a.Log.ErrorFatalF("could not create initial website configurations: %v", err)
	}

	if lib.GetDevelopmentMode() == "development" {
		listenAndServe(a.App, a.Log)
	} else {
		go listenAndServe(a.App, a.Log)
		startServerWithGracefulShutdown(a.App, a.Log)
	}
}

func listenAndServe(a *fiber.App, l *lib.AppLogger) {
	if err := a.Listen(lib.GetServerBindAddress()); err != nil {
		l.ErrorFatalF("Oops... Server is not running! Reason: %v\n", err)
	}
}

func startServerWithGracefulShutdown(a *fiber.App, l *lib.AppLogger) {
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
