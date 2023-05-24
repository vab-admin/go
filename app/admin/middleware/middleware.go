package middleware

import (
	"github.com/casbin/casbin/v2/util"
	"github.com/labstack/echo/v5"
)

type SkipperFunc func(echo.Context) bool

func SkipHandler(c echo.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}

// AllowRoutesSkipper
// @date 2023-05-11 20:58:39
func AllowRoutesSkipper(routes ...echo.Route) SkipperFunc {
	return func(c echo.Context) bool {
		for _, route := range routes {
			if util.KeyMatch2(c.Request().URL.String(), route.Path) && c.Request().Method == route.Method {
				return true
			}
		}

		return false
	}
}
