package middleware

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/ipfs/boxo/path"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

func SessionID(c echo.Context) string {
	return readCookie(c, "session")
}

var Cache = cacheHandler{}

type cacheHandler struct{}

func (c cacheHandler) GetChallenge(e echo.Context) protocol.URLEncodedBase64 {
	key := cacheKey(e, "challenge")
	if x, found := challengeStore.Get(key); found {
		chalbz := x.([]byte)
		chal := protocol.URLEncodedBase64{}
		err := chal.UnmarshalJSON(chalbz)
		if err != nil {
			return chal
		}
		return chal
	}
	chal, _ := protocol.CreateChallenge()
	// Save challenge
	str, err := chal.MarshalJSON()
	challengeStore.Set(key, str, cache.DefaultExpiration)
	if err != nil {
		return chal
	}
	return chal
}

func (c cacheHandler) GetLocalPath(e echo.Context) string {
	return ""
}

func (c cacheHandler) GetRemoteCID(e echo.Context) path.Path {
	return nil
}

func (c cacheHandler) SetLocalPath(e echo.Context, path string) {
}

func (c cacheHandler) SetRemoteCID(e echo.Context, path path.Path) {
}
