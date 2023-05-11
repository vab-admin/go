package middleware

import (
	"github.com/labstack/echo/v5"
	log "github.com/sirupsen/logrus"
	"strconv"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/contextx"
	"vab-admin/go/pkg/errors"
)

func Casbin(enforcer *service.Enforcer, skippers ...SkipperFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if SkipHandler(c, skippers...) {
				return next(c)
			}

			p := c.Path()
			m := c.Request().Method

			userID := contextx.FromUserID(c.Request().Context())

			ok, err := enforcer.Enforce(strconv.FormatUint(userID, 10), p, m)
			if err != nil {
				log.WithError(err).WithFields(map[string]any{
					"path":   p,
					"method": m,
					"userId": userID,
				}).Error("rbac权限验证失败")
				return errors.ErrInternalServer
			}

			if !ok {
				log.WithFields(map[string]any{
					"path":   p,
					"method": m,
					"userId": userID,
				}).Warn("无权限操作")
				return errors.New("无权限操作")
			}

			return next(c)
		}
	}
}
