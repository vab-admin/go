// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"erp/app/admin/handler"
	"erp/app/admin/repository"
	"erp/app/admin/router"
	"erp/app/admin/service"
)

// Injectors from wire.go:

// CreateApp
// @date 2023-05-07 19:04:57
func CreateApp() (*Injector, error) {
	adminUser := &repository.AdminUser{}
	adminUserGroup := &repository.AdminUserGroup{}
	adminGroup := &repository.AdminGroup{}
	serviceAdminUserGroup := &service.AdminUserGroup{
		AdminUserGroupRepo: adminUserGroup,
		AdminGroupRepo:     adminGroup,
	}
	serviceAdminUser := &service.AdminUser{
		AdminUserRepo:         adminUser,
		AdminUserGroupService: serviceAdminUserGroup,
	}
	handlerAdminUser := handler.NewAdminUser(serviceAdminUser)
	adminRule := &repository.AdminRule{}
	serviceAdminRule := &service.AdminRule{
		AdminRuleRepo: adminRule,
	}
	handlerAdminRule := &handler.AdminRule{
		AdminRuleService: serviceAdminRule,
	}
	adminGroupRule := &repository.AdminGroupRule{}
	serviceAdminGroupRule := &service.AdminGroupRule{
		AdminRuleRepo:      adminRule,
		AdminGroupRuleRepo: adminGroupRule,
	}
	serviceAdminGroup := &service.AdminGroup{
		AdminGroupRepo:        adminGroup,
		AdminGroupRuleService: serviceAdminGroupRule,
	}
	handlerAdminGroup := &handler.AdminGroup{
		AdminGroupService: serviceAdminGroup,
	}
	adminRuleAction := &repository.AdminRuleAction{}
	serviceAdminRuleAction := &service.AdminRuleAction{
		AdminRuleActionRepo: adminRuleAction,
	}
	handlerAdminRuleAction := &handler.AdminRuleAction{
		AdminRuleActionService: serviceAdminRuleAction,
	}
	route := &router.Route{
		AdminUserHandler:       handlerAdminUser,
		AdminRuleHandler:       handlerAdminRule,
		AdminGroupHandler:      handlerAdminGroup,
		AdminRuleActionHandler: handlerAdminRuleAction,
	}
	echo := newApp(route)
	injector := &Injector{
		App: echo,
	}
	return injector, nil
}