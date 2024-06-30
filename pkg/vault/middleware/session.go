package middleware

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
)

func SessionID(c echo.Context) string {
	return readCookie(c, "session")
}

func Challenge(c echo.Context) protocol.URLEncodedBase64 {
	chal, _ := protocol.CreateChallenge()
	return protocol.URLEncodedBase64(chal)
}
