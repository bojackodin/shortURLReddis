package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (app *application) run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	shutdownErr := make(chan error, 1)

	app.infoLog.Printf("Starting server on %s", srv.Addr)

	go func() {
		shutdownErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-shutdownErr:
		return err
	case <-ctx.Done():
	}

	app.infoLog.Printf("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	app.infoLog.Printf("stopped server")

	return nil
}
