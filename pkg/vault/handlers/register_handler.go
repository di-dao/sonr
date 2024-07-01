package handlers

import (
	"fmt"

	pages "github.com/di-dao/sonr/pkg/vault/components"
	"github.com/di-dao/sonr/pkg/vault/middleware"
	"github.com/labstack/echo/v4"
)

var Register = registerHandler{}

type registerHandler struct{}

func (h registerHandler) Start(e echo.Context) error {
	return e.JSON(0, nil)
}

func (h registerHandler) Finish(e echo.Context) error {
	cred, err := middleware.ParseCredentialFromForm(e)
	if err != nil {
		return e.JSON(500, err.Error())
	}
	return e.JSON(200, fmt.Sprintf("REGISTER: %s", string(cred.ID)))
}

func (h registerHandler) Page(e echo.Context) error {
	return middleware.Render(e, pages.Register(middleware.SessionID(e), string(middleware.Cache.GetChallenge(e))))
}
