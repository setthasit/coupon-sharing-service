package main

import (
	"context"
	"coupon-service/di"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := di.InitializeApp()

	app.Ctrl.RegisterRoute(app.Engine)

	srv := &http.Server{
		Addr:    ":9999",
		Handler: app.Engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server force to shutdown!: %q", err)
	}

	log.Println("Server gracefully shut down")
}
