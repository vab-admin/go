package db

import (
	"context"
	"erp/pkg/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
)

var redisDB *redis.Client

// NewRedis
// @param conf
// @date 2022-09-10 17:35:01
func NewRedis(conf config.Redis) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.Database,
	})

	if err := redisDB.Ping(context.Background()).Err(); err != nil {
		log.Fatal("连接redis失败", zap.String("error", err.Error()))
	}
}

// RedisDB
// @date 2022-09-10 17:35:01
func RedisDB() *redis.Client {
	return redisDB
}
