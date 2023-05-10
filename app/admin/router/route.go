package router

import (
	"erp/app/admin/handler"
	"erp/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"net/http"
)

type ErrHandle interface {
}

type Route struct {
	errorsMap map[string]echo.HandlerFunc `wire:"-"`

	AdminUserHandler       *handler.AdminUser
	AdminRuleHandler       *handler.AdminRule
	AdminGroupHandler      *handler.AdminGroup
	AdminRuleActionHandler *handler.AdminRuleAction
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
		//func(next echo.HandlerFunc) echo.HandlerFunc {
		//	return func(c echo.Context) error {
		//		time.Sleep(time.Second * 2)
		//		return next(c)
		//	}
		//},
	)

	v1Pub := app.Group("/api/v1")

	v1Auth := app.Group("/api/v1")

	r.Api(v1Auth)

	v1Pub.POST("/login", r.AdminUserHandler.Login)
	v1Pub.GET("/userInfo", r.AdminUserHandler.Info)

	v1Pub.POST("/admin/users", r.AdminUserHandler.Create)
	v1Pub.GET("/admin/users", r.AdminUserHandler.Query)
	v1Pub.GET("/admin/users/:id/edit", r.AdminUserHandler.Edit)
	v1Pub.DELETE("/admin/users/:id", r.AdminUserHandler.Delete)
	v1Pub.PUT("/admin/users/:id", r.AdminUserHandler.Update)

	v1Pub.GET("/admin/rules/tree", r.AdminRuleHandler.Tree)
	v1Pub.POST("/admin/rules", r.AdminRuleHandler.Create)
	v1Pub.GET("/admin/rules/:id/edit", r.AdminRuleHandler.Edit)
	v1Pub.PUT("/admin/rules/:id", r.AdminRuleHandler.Update)
	v1Pub.DELETE("/admin/rules/:id", r.AdminRuleHandler.Delete)
	v1Pub.PUT("/admin/rules/:id/field", r.AdminRuleHandler.UpdateField)

	v1Pub.GET("/admin/rules/:ruleId/actions", r.AdminRuleActionHandler.Query)
	v1Pub.POST("/admin/rules/:ruleId/actions", r.AdminRuleActionHandler.Create)
	v1Pub.GET("/admin/rules/actions/:id/edit", r.AdminRuleActionHandler.Edit)
	v1Pub.PUT("/admin/rules/:ruleId/actions/:id", r.AdminRuleActionHandler.Update)
	v1Pub.DELETE("/admin/rules/actions/:id", r.AdminRuleHandler.Delete)

	v1Pub.GET("/admin/group", r.AdminGroupHandler.Query)
	v1Pub.POST("/admin/group", r.AdminGroupHandler.Create)
	v1Pub.GET("/admin/group/:id/edit", r.AdminGroupHandler.Edit)
	v1Pub.PUT("/admin/group/:id", r.AdminGroupHandler.Update)
	v1Pub.DELETE("/admin/group/:id", r.AdminGroupHandler.Delete)
}
