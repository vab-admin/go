package middleware

import (
	"github.com/casbin/casbin/v2/util"
	"github.com/labstack/echo/v5"
	"github.com/thoas/go-funk"
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

// AllowPathSkipper
// @param paths
// @date 2023-05-11 20:44:00
func AllowPathSkipper(paths ...string) SkipperFunc {
	return func(c echo.Context) bool {
		path := c.Request().URL.Path
		return funk.InStrings(paths, path)
	}
}

// AllowPathWithMethodSkipper
// @param path
// @param method
// @date 2023-05-11 20:58:39
func AllowPathWithMethodSkipper(path, method string) SkipperFunc {
	return func(c echo.Context) bool {
		return util.KeyMatch2(c.Request().URL.String(), path) && c.Request().Method == method
	}
}
