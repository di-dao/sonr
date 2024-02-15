package highway

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	_ "github.com/sonrhq/sonr/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/sonrhq/sonr/pkg/highway/middleware"
	"github.com/sonrhq/sonr/pkg/highway/routes"
)

// Start starts the highway server
func Start(opts ...HighwayOption) {
	// Set the default options
	cnfg := NewHighwayOptions()
	for _, opt := range opts {
		opt(cnfg)
	}

	// Create the router
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	middleware.UseDefaults(e)
	routes.RegisterCosmosAPI(e)
	routes.RegisterSonrAPI(e)
	routes.RegisterHTMXPages(e)

	// Serve the router
	cnfg.PrintBanner()
	e.Logger.Fatal(e.Start(cnfg.listenAddress()))
}
