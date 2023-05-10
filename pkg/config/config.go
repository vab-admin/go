package config

import (
	"os"

	"go.uber.org/zap"

	"log"

	"github.com/spf13/viper"
)

type Database struct {
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Port        string `mapstructure:"port"`
	DbName      string `mapstructure:"dbname"`
	Migrate     bool   `mapstructure:"migrate"`
	Debug       bool   `mapstructure:"debug"`
	Driver      string `mapstructure:"driver"`
	MaxOpen     int    `mapstructure:"max_open"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxLifeTime string `mapstructure:"max_life_time"`
	MaxIdleTime string `mapstructure:"max_idle_time"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type JwtAuth struct {
	Secret string `mapstructure:"secret"`
}

type Config struct {
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	Oss      Oss      `mapstructure:"oss"`
	Casbin   Casbin   `mapstructure:"casbin"`
	Auth     JwtAuth  `mapstructure:"auth"`
}

type Oss struct {
	Domain     string `mapstructure:"domain" json:"domain"`
	AccessKey  string `mapstructure:"accessKey" json:"accessKey"`
	BucketName string `mapstructure:"bucketName" json:"bucketName"`
	Secret     string `mapstructure:"secret" json:"secret"`
	Endpoint   string `mapstructure:"endpoint" json:"endpoint"`
	Bucket     string `mapstructure:"bucket" json:"bucket"`
	Dir        string ` mapstructure:"dir" json:"dir"`
	HostName   string ` mapstructure:"hostName" json:"hostName"`
}

type Casbin struct {
	Log    bool
	RootId uint64 `mapstructure:"rootId"`
	Model  string `mapstructure:"model"`
}

var DefaultConfig Config

// NewConfig
// @date 2022-09-10 17:34:04
func NewConfig() Config {
	var (
		configPath = "./config"
		configName = "app"
	)

	if v := os.Getenv("CONFIG_PATH"); v != "" {
		configPath = v
	}

	if v := os.Getenv("CONFIG_NAME"); v != "" {
		configName = v
	}

	viper.SetConfigName(configName)

	viper.AddConfigPath(configPath)

	viper.AddConfigPath("$CONFIG")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigMarshalError:
			log.Fatal("配置文件解码失败")
		case viper.ConfigFileNotFoundError:
			log.Fatal("配置文件不存在", zap.String("name", configName), zap.String("path", configPath))
		default:
			log.Fatal("读取配置文件失败", zap.String("name", configName), zap.String("path", configPath))
		}
	}

	if err := viper.Unmarshal(&DefaultConfig); err != nil {
		log.Fatal("解码失败", zap.String("name", configName), zap.String("path", configPath))
	}

	return DefaultConfig
}

func IsDev() bool {
	return Mode == "dev"
}
