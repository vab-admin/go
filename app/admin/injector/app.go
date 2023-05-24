package injector

import (
	"github.com/labstack/echo/v5"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/app/admin/router"
	"vab-admin/go/pkg/json_serializer"
)

type Logger struct{}

func (l *Logger) Write(p []byte) (n int, err error) {
	log.Info(string(p))

	return 0, nil
}

func (l *Logger) Error(err error) {
	log.Error(err.Error())
}

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
	app.Logger = &Logger{}

	router.RegisterHandlers(app)

	return app
}
