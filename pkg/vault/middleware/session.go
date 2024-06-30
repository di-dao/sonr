package middleware

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
)

func SessionID(c echo.Context) string {
	return readCookie(c, "session")
}

func Challenge(c echo.Context) protocol.URLEncodedBase64 {
	val, err := getCacheKV(c, "challenge")
	if err != nil {
		chal, _ := protocol.CreateChallenge()
		setCacheKV(c, "challenge", chal.String())
	}
	return protocol.URLEncodedBase64(val)
}
