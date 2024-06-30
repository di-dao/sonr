package vault

import (
	"context"
	"github.com/di-dao/sonr/pkg/vault/middleware"
	"github.com/di-dao/sonr/pkg/vault/routes"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Serve(ctx context.Context) {
	e := echo.New()
	middleware.UseCache(e)
	routes.RegisterPages(e)
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
