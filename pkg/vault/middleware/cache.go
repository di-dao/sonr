package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

var (
	challengeStore = cache.New(5*time.Minute, 10*time.Minute)
	localFSStore   = cache.New(30*time.Minute, 1*time.Hour)
	remoteFSStore  = cache.New(30*time.Minute, 1*time.Hour)
)

func cacheKey(e echo.Context, key string) string {
	return fmt.Sprintf(SessionID(e), ".", key)
}
