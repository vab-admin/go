package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/app/admin/injector"
	"vab-admin/go/pkg/config"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/model"
)

// initLogger
// @date 2023-05-11 00:51:13
func initLogger() {
	formatter := &log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}

	log.SetLevel(log.TraceLevel)
	log.SetFormatter(formatter)

}

// createDatabase
// @date 2023-05-11 00:51:16
func createDatabase(conf config.Config) {

	db.NewDB(conf.Database)
	db.NewRedis(conf.Redis)

	_ = db.Instance(context.Background()).AutoMigrate(
		&model.AdminUser{},
		&model.AdminRule{},
		&model.AdminGroup{},
		&model.SystemApi{},
	)
}

// createApp
// @date 2023-05-11 00:51:21
func createApp(conf config.Config) {

	app, err := injector.CreateApp(conf)
	if err != nil {
		log.WithError(err).Fatal("创建应用失败")
		return
	}

	app.Run(":8080")
}

func main() {
	initLogger()
	conf := config.NewConfig()

	createDatabase(conf)
	createApp(conf)
}
