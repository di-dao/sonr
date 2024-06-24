package handlers

import (
	"github.com/di-dao/sonr/pkg/vault/middleware"
	"github.com/di-dao/sonr/pkg/vault/pages"
	"github.com/labstack/echo/v4"
)

func HandleRegisterPage(e echo.Context) error {
	return middleware.Render(e, pages.Home())
}
