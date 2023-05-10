package db

import (
	"context"
	"database/sql"
	"erp/pkg/config"
	"erp/pkg/contextx"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var (
	db *gorm.DB
)

// NewDB
// @param conf
// @date 2022-09-10 17:34:41
func NewDB(conf config.Database) {
	db = newDb(conf)
}

func CreateDb(conf config.Database) *gorm.DB {
	return newDb(conf)
}

// newDb
// @param conf
// @date 2022-09-10 17:34:39
func newDb(conf config.Database) *gorm.DB {
	ormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		PrepareStmt:                              true,
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if conf.Debug {
		ormConfig.Logger = ormConfig.Logger.LogMode(logger.Info)
	}

	var dsn string
	var conn *gorm.DB
	var err error

	switch conf.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
		conn, err = gorm.Open(mysql.Open(dsn), ormConfig)
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", conf.Host, conf.User, conf.Password, conf.Port, conf.DbName)
		conn, err = gorm.Open(postgres.Open(dsn), ormConfig)
	case "mssql":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	default:
		dsn = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", conf.Host, conf.User, conf.Password, conf.Port, conf.DbName)
		conn, err = gorm.Open(postgres.Open(dsn), ormConfig)
	}

	if err != nil {
		log.Fatal("连接数据库失败")
	}

	var tx *sql.DB
	tx, err = conn.DB()
	if err != nil {
		log.Fatal("获取数据库失败", zap.String("error", err.Error()))
	}

	// 设置最大的并发打开连接数
	// 设置这个数小于等于0则表示没有显示
	if conf.MaxOpen > 0 && tx != nil {
		tx.SetMaxOpenConns(conf.MaxOpen)
	}

	// 设置最大的空闲连接数
	// 设置小于等于0的数意味着不保留空闲连接
	if conf.MaxIdle > 0 && tx != nil {
		tx.SetMaxIdleConns(conf.MaxIdle)
	}

	if conf.MaxLifeTime != "" && tx != nil {
		var maxLifeTime time.Duration
		if maxLifeTime, err = time.ParseDuration(conf.MaxLifeTime); err == nil {
			// 设置连接的最大生命周期
			// 设置为0的话意味着没有最大生命周期，连接总是可重用(默认行为)。
			tx.SetConnMaxLifetime(maxLifeTime)
		}
	}

	if conf.MaxIdleTime != "" && tx != nil {
		var maxIdleTime time.Duration
		if maxIdleTime, err = time.ParseDuration(conf.MaxIdleTime); err == nil {
			// 连接最大空闲时间
			tx.SetConnMaxIdleTime(maxIdleTime)
		}
	}

	return conn
}

// Session
// @param ctx
// @date 2022-09-10 17:34:38
func Session(ctx context.Context) *gorm.DB {
	trans, has := contextx.FromTrans(ctx)
	if !has {
		return db.WithContext(ctx)
	}

	tx, ok := trans.(*gorm.DB)
	if ok {
		return tx
	}

	return db.WithContext(ctx)
}

// Transaction
// @param ctx
// @param fc
// @date 2022-09-10 17:34:37
func Transaction(ctx context.Context, fc func(context.Context) error) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		v := contextx.NewTrans(ctx, tx)
		return fc(v)
	})
}

func Instance(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
