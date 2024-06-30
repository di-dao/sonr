package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coocood/freecache"
	"github.com/labstack/echo/v4"
	"gopkg.in/mcuadros/go-defaults.v1"
)

// Config defiens the configuration for a cache middleware.
type CacheConfig struct {
	// TTL time to life of the cache.
	TTL time.Duration `default:"1m"`
	// Methods methods to be cached.
	Methods []string `default:"[GET]"`
	// StatusCode method to be cached.
	StatusCode []int `default:"[200,404]"`
	// IgnoreQuery if true the Query values from the requests are ignored on
	// the key generation.
	IgnoreQuery bool
	// Refresh fuction called before use the cache, if true, the cache is deleted.
	Refresh func(r *http.Request) bool
	// Cache fuction called before cache a request, if false, the request is not
	// cached. If set Method is ignored.
	Cache func(r *http.Request) bool
}

type CacheMiddleware struct {
	cfg   *CacheConfig
	cache *freecache.Cache
}

func RequestCaching(cache *freecache.Cache) echo.MiddlewareFunc {
	cfg := &CacheConfig{}

	defaults.SetDefaults(cfg)
	m := &CacheMiddleware{cfg: cfg, cache: cache}
	return m.Handler
}

func (m *CacheMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !m.isCacheable(c.Request()) {
			return next(c)
		}

		if mayHasBody(c.Request().Method) {
			c.Logger().Warnf("request with body are cached ignoring the content")
		}

		key := m.getKey(c.Request())
		err := m.readCache(key, c)
		if err == nil {
			return nil
		}

		if err != freecache.ErrNotFound {
			c.Logger().Errorf("error reading cache: %s", err)
		}

		recorder := NewResponseRecorder(c.Response().Writer)
		c.Response().Writer = recorder

		err = next(c)
		if err := m.cacheResult(key, recorder); err != nil {
			c.Logger().Error(err)
		}

		return err
	}
}

func (m *CacheMiddleware) readCache(key []byte, c echo.Context) error {
	if m.cfg.Refresh != nil && m.cfg.Refresh(c.Request()) {
		return freecache.ErrNotFound
	}

	value, err := m.cache.Get(key)
	if err != nil {
		return err
	}

	entry := &CacheEntry{}
	if err := entry.Decode(value); err != nil {
		return err
	}

	return entry.Replay(c.Response())
}

func (m *CacheMiddleware) cacheResult(key []byte, r *ResponseRecorder) error {
	e := r.Result()
	b, err := e.Encode()
	if err != nil {
		return fmt.Errorf("unable to read recorded response: %s", err)
	}

	if !m.isStatusCacheable(e) {
		return nil
	}

	return m.cache.Set(key, b, int(m.cfg.TTL.Seconds()))
}

func (m *CacheMiddleware) isStatusCacheable(e *CacheEntry) bool {
	for _, status := range m.cfg.StatusCode {
		if e.StatusCode == status {
			return true
		}
	}

	return false
}

func (m *CacheMiddleware) isCacheable(r *http.Request) bool {
	if m.cfg.Cache != nil {
		return m.cfg.Cache(r)
	}

	for _, method := range m.cfg.Methods {
		if r.Method == method {
			return true
		}
	}

	return false
}

func (m *CacheMiddleware) getKey(r *http.Request) []byte {
	base := r.Method + "|" + r.URL.Path
	if !m.cfg.IgnoreQuery {
		base += "|" + r.URL.Query().Encode()
	}

	return []byte(base)
}

func mayHasBody(method string) bool {
	m := method
	if m == http.MethodPost || m == http.MethodPut || m == http.MethodDelete || m == http.MethodPatch {
		return true
	}

	return false
}
