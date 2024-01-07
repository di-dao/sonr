package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pterm/pterm"

	"github.com/sonrhq/sonr/gateway/routes"
)

func Start() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Mount("/", routes.HomeRoutes())
    pterm.DefaultHeader.Println("Running Gateway: http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
