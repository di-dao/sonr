package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
	"github.com/EchoVault/EchoVault/pkg/echovault"
)

var cache *echovault.EchoVault

func init() {
	var err error
	cache, err = echovault.NewEchoVault()
	if err != nil {
		panic(err)
	}
}

func Cache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := SessionID(c)
		if sessionID == "" {
			sessionID = ksuid.New().String()
			writeCookie(c, "session", sessionID)
		}

		cacheKey := sessionID + ":" + c.Request().URL.Path

		// Check if response is cached
		cachedResponse, err := cache.Get(cacheKey)
		if err == nil {
			// If cached, return the cached response
			cacheEntry := &CacheEntry{}
			err = cacheEntry.Decode([]byte(cachedResponse))
			if err == nil {
				return cacheEntry.Replay(c.Response())
			}
		}

		// If not cached, process the request
		recorder := NewResponseRecorder(c.Response().Writer)
		c.Response().Writer = recorder

		err = next(c)
		if err != nil {
			return err
		}

		// Cache the response
		cacheEntry := recorder.Result()
		encodedEntry, err := cacheEntry.Encode()
		if err == nil {
			cache.Set(cacheKey, string(encodedEntry), echovault.SETOptions{})
		}

		return nil
	}
}
