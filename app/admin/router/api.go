package router

import (
	"github.com/labstack/echo/v5"
	"vab-admin/go/app/admin/middleware"
)

// Api
// @param api
// @date 2023-05-11 00:40:12
func (r *Route) Api(api *echo.Group) {

	api.GET("/userInfo", r.AdminUserHandler.Info)
	api.GET("/router", r.AdminUserHandler.Router)

	adminRoute := api.Group("/admin", middleware.Casbin(r.Enforcer, nil))
	{
		apiRouters("管理员用户", "/users", adminRoute, r.AdminUserHandler)

		apiRouters("菜单规则", "/rules", adminRoute, r.AdminRuleHandler)
		adminRoute.PUT("/rules/:id/field", r.AdminRuleHandler.UpdateField) // 更新菜单规则字段

		apiRouters("系统接口", "/api", adminRoute, r.SystemApiHandler)
		adminRoute.GET("/api/all", r.SystemApiHandler.All) // 查询按钮接口

		apiRouters("管理员角色", "/roles", adminRoute, r.AdminRoleHandler)
	}
}
