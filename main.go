package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func debug(x interface{}) {
	fmt.Printf("%T: %#v\n", x, x)
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/health_check", func(ctx iris.Context) {
		ctx.StatusCode(200)
		ctx.WriteString("Ok")
	})

	app.Run(iris.Addr(":8080"))
}
