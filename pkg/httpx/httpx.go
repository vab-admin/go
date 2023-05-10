package httpx

import (
	"github.com/labstack/echo/v5"
	log "github.com/sirupsen/logrus"
	"net/http"

	"erp/pkg/errors"
)

// Bind
// @param c
// @param v
// @date 2022-07-07 23:52:18
func Bind(c echo.Context, v interface{}) error {
	if err := c.Bind(v); err != nil {
		log.WithError(err).Warn("解析参数失败")
		return errors.ErrInvalidParam
	}
	return nil
}

// OK
// @param c
// @param msg
// @date 2022-07-07 23:52:17
func OK(c echo.Context, msg string) error {
	return c.JSON(http.StatusOK, errors.Response{Status: http.StatusOK, Code: errors.CodeOK, Msg: msg})
}

// OkJSON
// @param c
// @param msg
// @param data
// @date 2022-07-07 23:52:05
func OkJSON(c echo.Context, msg string, data any) error {
	return c.JSON(http.StatusOK, errors.Response{Status: http.StatusOK, Code: errors.CodeOK, Msg: msg, Data: data})
}
