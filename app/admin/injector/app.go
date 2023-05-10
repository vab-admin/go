package injector

import (
	"github.com/labstack/echo/v5"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/app/admin/router"
	"vab-admin/go/pkg/json_serializer"
)

// Run
// @date: 2022-02-01 18:48:03
func (i *Injector) Run(address string) {

	if err := i.App.Start(address); err != nil {
		log.WithError(err).Error("服务器启动失败")
	}

}

// newApp
// @date: 2022-02-01 18:48:08
func newApp(router *router.Route) *echo.Echo {
	var app = echo.New()

	app.JSONSerializer = &json_serializer.JsonSerializer{}

	router.RegisterHandlers(app)

	return app
}
