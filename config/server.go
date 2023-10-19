package config

import "github.com/kataras/iris/v12"

var Config = createAppConfig()

func (c *AppConfig) initilize(port string) {
	c.app = iris.New()
	c.port = port
	c.app.Use(iris.Compression)
	c.app.Listen(":" + c.port)
}

func (c *AppConfig) StartServer(port string) {
	c.initilize(port)
}