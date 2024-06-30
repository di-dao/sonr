package handlers

import (
	pages "github.com/di-dao/sonr/pkg/vault/components"
	"github.com/di-dao/sonr/pkg/vault/middleware"
	"github.com/labstack/echo/v4"
)

var Register = registerHandler{}

type registerHandler struct{}

func (h registerHandler) Page(e echo.Context) error {
	return middleware.Render(e, pages.Register(middleware.SessionID(e), string(middleware.Cache.GetChallenge(e))))
}
