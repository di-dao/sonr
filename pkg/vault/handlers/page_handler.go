package handlers

import (
	"github.com/di-dao/sonr/pkg/vault/components"
	"github.com/di-dao/sonr/pkg/vault/middleware"
	"github.com/labstack/echo/v4"
)

var Session = sessionHandler{}

type sessionHandler struct{}

func (h sessionHandler) Page(e echo.Context) error {
	return middleware.Render(e, components.Home(middleware.SessionID(e)))
}
