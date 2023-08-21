package lib

import (
	"log"
	"os"
)

type AppLogger struct {
	logger      *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger() *AppLogger {
	prefix := "[soramon0/portfolio] "
	flags := log.LstdFlags
	stdLogger := log.New(os.Stdout, prefix, flags)
	infoLogger := log.New(os.Stdout, prefix+"INFO: ", flags)
	errorLogger := log.New(os.Stdout, prefix+"ERROR: ", flags)

	return &AppLogger{
		logger:      stdLogger,
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}
}

func (l *AppLogger) Info(v ...any) {
	l.infoLogger.Print(v...)
}

func (l *AppLogger) Infof(format string, v ...any) {
	l.infoLogger.Printf(format, v...)
}

func (l *AppLogger) Error(v ...any) {
	l.errorLogger.Print(v...)
}

func (l *AppLogger) ErrorF(format string, v ...any) {
	l.errorLogger.Printf(format, v...)
}

func (l *AppLogger) ErrorFatal(v ...any) {
	l.errorLogger.Fatal(v...)
}

func (l *AppLogger) ErrorFatalF(format string, v ...any) {
	l.errorLogger.Fatalf(format, v...)
}
