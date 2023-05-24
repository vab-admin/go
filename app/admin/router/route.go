package router

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"vab-admin/go/app/admin/handler"
	middleware2 "vab-admin/go/app/admin/middleware"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/errors"
)

type Route struct {
	errorsMap map[string]echo.HandlerFunc `wire:"-"`
	Enforcer  *service.Enforcer

	AdminUserHandler *handler.AdminUser
	AdminRuleHandler *handler.AdminRule
	AdminRoleHandler *handler.AdminRole
	SystemApiHandler *handler.SystemApi
}

type IRoute interface {
	Query(ctx echo.Context) error
	Create(ctx echo.Context) error
	Edit(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
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
	v1Auth := app.Group("/api/v1",
		middleware2.JwtAuth(),
	)

	r.Api(v1Auth)

	v1Pub.POST("/login", r.AdminUserHandler.Login)

}

func apiRouters(name, prefix string, api *echo.Group, handler IRoute) {

	routers := []echo.Route{
		{Method: http.MethodGet, Path: prefix, Handler: handler.Query, Name: "查询" + name},
		{Method: http.MethodPost, Path: prefix, Handler: handler.Create, Name: "创建" + name},
		{Method: http.MethodGet, Path: prefix + "/:id/edit", Handler: handler.Edit, Name: "编辑" + name},
		{Method: http.MethodPut, Path: prefix + "/:id", Handler: handler.Update, Name: "更新" + name},
		{Method: http.MethodDelete, Path: prefix + "/:id", Handler: handler.Delete, Name: "删除" + name},
	}

	for _, router := range routers {
		if _, err := api.AddRoute(router); err != nil {
			log.WithError(err).Fatal("添加路由失败")
		}
	}
}
