package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"net/http"
	"vab-admin/go/pkg/auth/jwtauth"
	"vab-admin/go/pkg/contextx"
	"vab-admin/go/pkg/errors"
)

var jwtConf = middleware.JWTConfig{
	ContextKey:   "user",
	TokenLookup:  "query:token,header:token,header:Authorization",
	ErrorHandler: jwtErrorHandler,
	ParseTokenFunc: func(c echo.Context, auth string, source middleware.ExtractorSource) (interface{}, error) {
		claims := jwtauth.UserInfoClaims{}

		_, err := jwt.ParseWithClaims(auth, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("api"), nil
		}, jwt.WithValidMethods([]string{"HS256"}))
		if err != nil {
			return nil, err
		}

		ctx := c.Request().Context()
		ctx = contextx.NewUserId(ctx, claims.UserID)

		c.SetRequest(c.Request().WithContext(ctx))
		return claims, err
	},
}

var (
	ErrUnauthorized = &errors.Response{Status: http.StatusUnauthorized, Code: http.StatusUnauthorized, Msg: "登录无效或已过期"}
)

func JwtAuth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(jwtConf)
}

func jwtErrorHandler(c echo.Context, _ error) error {
	return c.JSON(http.StatusUnauthorized, ErrUnauthorized)
}
