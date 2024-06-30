package middleware

import (
	"fmt"

	"github.com/echovault/echovault/echovault"
	"github.com/labstack/echo/v4"
)

var cache *echovault.EchoVault

func init() {
	var err error
	cache, err = echovault.NewEchoVault()
	if err != nil {
		panic(err)
	}
}

func setCacheKV(e echo.Context, key string, value string) error {
	_, _, err := cache.Set(cacheKey(e, key), value, echovault.SetOptions{})
	return err
}

func getCacheKV(e echo.Context, key string) (string, error) {
	return cache.Get(cacheKey(e, key))
}

func cacheKey(e echo.Context, key string) string {
	return fmt.Sprintf(SessionID(e), ".", key)
}
