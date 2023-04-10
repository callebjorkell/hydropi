package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05",
		FullTimestamp:   true,
	})

	if err := RootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

func startServer() {
	log.Info("Starting server")
	ctx, cancel := ContextWithCancelOnSignal()
	defer cancel()

	// TODO: server stuff here.

	<-ctx.Done()
	log.Info("Server gracefully shut down")
}

// ContextWithCancelOnSignal creates a context that has an explicit cancel, as well as a cancel if a SIGTERM or SIGINT
// is received by the application.
func ContextWithCancelOnSignal() (context.Context, context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-signalChan
		cancel()
	}()
	return ctx, cancel
}
