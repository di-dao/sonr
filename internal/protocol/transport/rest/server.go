package rest

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/helmet/v2"
)

type HttpTransport struct {
	*fiber.App
	SessionStore *session.Store
	ClientContext   client.Context
}

func NewHttpTransport(ctx client.Context) *HttpTransport {
	// HttPTransport
	rest := &HttpTransport{
		App: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
		ClientContext: ctx,
		SessionStore: session.New(),
	}

	// Middleware
	rest.Use(cors.New())
	rest.Use(helmet.New())

	// Status Methods
	rest.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK. Highway version v0.6.0. Running on HTTP/TLS")
	})

	// Query Methods
	rest.Get("/highway/query/service/:origin/:username", timeout.New(rest.QueryService, time.Second*5))
	rest.Get("/highway/query/document/:did", timeout.New(rest.QueryDocument, time.Second*5))

	// Auth Methods
	rest.Post("/highway/auth/keygen", timeout.New(rest.Keygen, time.Second*10))
	rest.Post("/highway/auth/login", timeout.New(rest.Login, time.Second*10))
	rest.Post("/highway/vault/add", timeout.New(rest.AddShare, time.Second*5))
	rest.Post("/highway/vault/sync", timeout.New(rest.SyncShare, time.Second*5))
	return rest
}