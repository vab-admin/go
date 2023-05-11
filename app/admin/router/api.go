package router

import "github.com/labstack/echo/v5"

// Api
// @param api
// @date 2023-05-11 00:40:12
func (r *Route) Api(api *echo.Group) {

	api.GET("/userInfo", r.AdminUserHandler.Info)

	adminRoute := api.Group("/admin")
	{
		adminRoute.GET("/users", r.AdminUserHandler.Query)         // 查询管理员用户
		adminRoute.POST("/users", r.AdminUserHandler.Create)       // 创建管理员用户
		adminRoute.GET("/:id/edit", r.AdminUserHandler.Edit)       // 编辑管理员用户
		adminRoute.DELETE("/users/:id", r.AdminUserHandler.Delete) // 删除管理员用户
		adminRoute.PUT("/users/:id", r.AdminUserHandler.Update)    // 更新管理员用户

		adminRoute.GET("/rules/tree", r.AdminRuleHandler.Tree)             // 获取规则树
		adminRoute.POST("/rules", r.AdminRuleHandler.Create)               // 创建菜单规则
		adminRoute.GET("/rules/:id/edit", r.AdminRuleHandler.Edit)         //编辑菜单规则
		adminRoute.PUT("/rules/:id", r.AdminRuleHandler.Update)            // 更新菜单规则
		adminRoute.DELETE("/rules/:id", r.AdminRuleHandler.Delete)         // 删除菜单规则
		adminRoute.PUT("/rules/:id/field", r.AdminRuleHandler.UpdateField) // 更新菜单规则字段

		adminRoute.GET("/api/all", r.SystemApiHandler.All)       // 查询按钮接口
		adminRoute.GET("/api", r.SystemApiHandler.Query)         // 查询按钮接口
		adminRoute.POST("/api", r.SystemApiHandler.Create)       // 创建按钮接口
		adminRoute.GET("/api/:id/edit", r.SystemApiHandler.Edit) // 编辑按钮接口
		adminRoute.PUT("/api/:id", r.SystemApiHandler.Update)    // 更新按钮接口
		adminRoute.DELETE("/api/:id", r.AdminRuleHandler.Delete) // 删除按钮接口

		adminRoute.GET("/group", r.AdminGroupHandler.Query)         // 查询管理员分组
		adminRoute.POST("/group", r.AdminGroupHandler.Create)       // 创建管理员分组
		adminRoute.GET("/group/:id/edit", r.AdminGroupHandler.Edit) // 编辑管理员分组
		adminRoute.PUT("/group/:id", r.AdminGroupHandler.Update)    // 更新管理员分组
		adminRoute.DELETE("/group/:id", r.AdminGroupHandler.Delete) // 删除管理员分组
	}
}
