package config

import "github.com/kataras/iris/v12"

type AppConfig struct {
	app *iris.Application
	port string
}

func createAppConfig() *AppConfig {
	return &AppConfig{}
}