package router

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"net/http"
	"vab-admin/go/app/admin/handler"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/errors"
)

type Route struct {
	errorsMap map[string]echo.HandlerFunc `wire:"-"`
	Enforcer  *service.Enforcer

	AdminUserHandler  *handler.AdminUser
	AdminRuleHandler  *handler.AdminRule
	AdminGroupHandler *handler.AdminGroup
	SystemApiHandler  *handler.SystemApi
}

// OnError
// @param err
// @param c
// @date 2022-06-04 19:40:30
func (r *Route) OnError(c echo.Context, err error) {

	next, ok := r.errorsMap[c.RouteInfo().Name()]
	if ok {
		err = next(c)
	}

	switch e := err.(type) {
	case *errors.Response:
		_ = c.JSON(e.Status, e)
	case validation.Errors:
		_ = c.JSON(http.StatusOK, errors.New(e.Error()))
	default:
		_ = c.JSON(http.StatusOK, errors.New(e.Error()))
	}
}

// OnNotFound
// @param c
// @date 2022-06-04 19:40:29
func OnNotFound(c echo.Context) error {
	return errors.ErrAPINotFound
}

// OnMethodNotAllowed
// @param c
// @date 2022-06-04 19:40:28
func OnMethodNotAllowed(c echo.Context) error {
	return errors.ErrAPINotFound
}

// RegisterHandlers
// @param app
// @date 2022-06-04 19:40:27
func (r *Route) RegisterHandlers(app *echo.Echo) {
	r.errorsMap = map[string]echo.HandlerFunc{
		echo.NotFoundRouteName:         OnNotFound,
		echo.MethodNotAllowedRouteName: OnMethodNotAllowed,
	}

	app.HTTPErrorHandler = r.OnError

	app.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.CORS(),
	)

	v1Pub := app.Group("/api/v1")
	v1Auth := app.Group("/api/v1") //middleware2.Casbin(r.Enforcer),

	r.Api(v1Auth)

	v1Pub.POST("/login", r.AdminUserHandler.Login)

}
