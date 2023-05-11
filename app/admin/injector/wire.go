//go:build wireinject

package injector

import (
	"github.com/google/wire"
	"vab-admin/go/app/admin/handler"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/app/admin/router"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/config"
)

// CreateApp
// @date 2023-05-07 19:04:57
func CreateApp(conf config.Config) (*Injector, error) {
	wire.Build(
		Set,
		newApp,
		service.NewCasbin,

		wire.NewSet(wire.Struct(new(router.Route), "*")),

		wire.NewSet(handler.NewAdminUser),
		wire.NewSet(wire.Struct(new(handler.AdminRule), "*")),
		wire.NewSet(wire.Struct(new(handler.AdminGroup), "*")),
		wire.NewSet(wire.Struct(new(handler.SystemApi), "*")),

		wire.NewSet(wire.Struct(new(service.AdminUser), "*")),
		wire.NewSet(wire.Struct(new(service.AdminRule), "*")),
		wire.NewSet(wire.Struct(new(service.AdminUserGroup), "*")),
		wire.NewSet(wire.Struct(new(service.AdminGroup), "*")),
		wire.NewSet(wire.Struct(new(service.AdminGroupRule), "*")),
		wire.NewSet(wire.Struct(new(service.SystemApi), "*")),
		wire.NewSet(wire.Struct(new(service.AdminRuleApi), "*")),

		wire.NewSet(wire.Struct(new(repository.AdminUser), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminRuleApi), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminUserGroup), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminRule), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminGroup), "*")),
		wire.NewSet(wire.Struct(new(repository.AdminGroupRule), "*")),
		wire.NewSet(wire.Struct(new(repository.SystemApi), "*")),
	)

	return new(Injector), nil
}
