package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/app/admin/injector"
	"vab-admin/go/pkg/config"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/model"
)

func initLogger() {
	formatter := &log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}

	log.SetLevel(log.TraceLevel)
	log.SetFormatter(formatter)

}

func createDatabase() {
	conf := config.NewConfig()

	db.NewDB(conf.Database)
	db.NewRedis(conf.Redis)

	_ = db.Instance(context.Background()).AutoMigrate(
		&model.AdminUser{},
		&model.AdminRule{},
		&model.AdminGroup{},
		&model.AdminRuleAction{},
	)
}

func createApp() {

	app, err := injector.CreateApp()
	if err != nil {
		log.WithError(err).Fatal("创建应用失败")
		return
	}

	app.Run(":8080")
}

func main() {
	initLogger()
	createDatabase()

	createApp()
}
