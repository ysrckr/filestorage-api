package main

import (
	"fmt"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/ysrckr/filestorage/api/config"
)




func main() {

	crs := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowCredentials: true,
})

	app := iris.New()
	app.Use(iris.Compression)
	app.UseRouter(crs)
	app.Get("/", func(ctx iris.Context) {
		cookie := ctx.GetCookie("go_cookie")
		if cookie == "" {
		ctx.SetCookie(&iris.Cookie{
			Name: "go_cookie",
			Value: "deger",
			HttpOnly: true,
			Secure: false,
		})
		}

		fmt.Println(cookie)


		ctx.JSON(iris.Map{
			"message": "Hello World!",
		})
	})
		

	app.Listen(":" + config.PORT)

	
}