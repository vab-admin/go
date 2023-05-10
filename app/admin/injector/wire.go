//go:build wireinject

package injector

import (
	"erp/app/admin/handler"
	"erp/app/admin/repository"
	"erp/app/admin/router"
	"erp/app/admin/service"
	"github.com/google/wire"
)

// CreateApp
// @date 2023-05-07 19:04:57
func CreateApp() (*Injector, error) {
	wire.Build(
		Set,
		newApp,

		wire.NewSet(wire.Struct(new(router.Route), "*")),

		wire.NewSet(handler.NewAdminUser),
		wire.NewSet(wire.Struct(new(handler.AdminRule), "*")),
		wire.NewSet(wire.Struct(new(handler.AdminGroup), "*")),
		wire.NewSet(wire.Struct(new(handler.AdminRuleAction), "*")),


		wire.NewSet(wire.Struct(new(service.AdminUser), "*")),
		wire.NewSet(wire.Struct(new(service.AdminRule), "*")),
		wire.NewSet(wire.Struct(new(service.AdminUserGroup), "*")),
		wire.NewSet(wire.Struct(new(service.AdminGroup), "*")),
		wire.NewSet(wire.Struct(new(service.AdminGroupRule), "*")),
		wire.NewSet(wire.Struct(new(service.AdminRuleAction), "*")),

		wire.NewSet(wire.Struct(new(repository.AdminUser), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminUserGroup), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminRule), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminGroup), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminGroupRule), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminRuleAction), "*")),
	)

	return new(Injector), nil
}
