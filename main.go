package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
  app.Use(iris.Compression)

	app.Listen(":8080")
}