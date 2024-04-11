package config

import (
	"os"

	"github.com/spf13/viper"
)

const (
	DefaultConfig = "default"
)

var (
	Env    = getEnv()
	Config = new()
	App    = Config.App
	Auth   = Config.Auth
	Mysql  = Config.Mysql
)

type config struct {
	App   app
	Auth  auth
	Mysql mysql
}

type app struct {
	Name     string
	Url      string
	Port     string
	Secret   string
	GinMode  string `mapstructure:"gin_mode"`
	LogLevel string `mapstructure:"log_level"`
}

type auth struct {
	JwtExpHours int `mapstructure:"jwt_exp_hours"`
}

type mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Options  string
}

func new() config {
	v := viper.New()
	v.AddConfigPath("./configs")

	// 先使用默认的配置文件，再覆盖对应环境的配置文件
	readConfig(v, DefaultConfig)
	readConfig(v, Env)

	// 最后再覆盖环境变量的配置
	v.AutomaticEnv()

	var c config
	err := v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	return c
}

func getEnv() string {
	env := os.Getenv("APP_ENV")

	switch env {
	case "production", "test":
		return env
	default:
		return "development"
	}
}

func readConfig(v *viper.Viper, file string) {
	v.SetConfigName(file)
	err := v.MergeInConfig()
	if err != nil {
		panic(err)
	}
}
