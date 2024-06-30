package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

func SessionCookies(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if val := readCookie(c, "session"); val == "" {
			writeCookie(c, "session", ksuid.New().String())
		}
		return next(c)
	}
}

func SessionID(c echo.Context) string {
	return readCookie(c, "session")
}

func readCookie(c echo.Context, key string) string {
	cookie, err := c.Cookie(key)
	if err != nil {
		return ""
	}
	if cookie == nil {
		return ""
	}
	return cookie.Value
}

func writeCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

// ClearSessionCache clears the cache for a specific session
func ClearSessionCache(c echo.Context) {
	sessionID := SessionID(c)
	if sessionID != "" {
		cache.Del(sessionID + ":*")
	}
}
