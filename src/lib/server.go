package lib

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App, l *AppLogger) {
	if GetDevelopmentMode() == "development" {
		listenAndServe(a, l)
	} else {
		go listenAndServe(a, l)
		startServerWithGracefulShutdown(a, l)
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
