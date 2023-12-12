package main

import (
	"github.com/kataras/iris/v12"
	"github.com/ysrckr/filestorage/api/config"
)




func main() {
	app := iris.New()
	app.Use(iris.Compression)
	app.Get("/", func(ctx iris.Context) {
		ctx.SetCookie(&iris.Cookie{
			Name: "go_cookie",
			Value: "deger",
			SameSite: iris.SameSiteLaxMode,
			HttpOnly: true,
			Secure: false,
		})

		ctx.JSON(iris.Map{
			"message": "Hello World!",
		})
	})
		

	app.Listen(":" + config.PORT)

	
}