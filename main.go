package main

import (
	"fmt"
	"log"

	"github.com/chenrihong/gostock/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()

	controllers.RegRouters(app)

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("stocks/today")
		return
	})

	Start(app)
}

func Start(app *iris.Application) {

	// This will serve the ./static/favicon.ico to: localhost:8080/favicon.ico
	app.Favicon("./static/favicon.ico")
	// that can recover from any http-relative panics
	app.Use(recover.New())
	app.AllowMethods(iris.MethodOptions)
	// file server
	app.HandleDir("/static", "./static")

	// RegisterView
	tmpl := iris.HTML("./views", ".html").Layout("shared/layout.html").Reload(true)

	// To add a template function use the AddFunc method of the preferred view engine.
	tmpl.AddFunc("Sum", func(i, j int64) int64 {
		return i + j
	})

	app.RegisterView(tmpl)
	// handle error
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		ctx.ViewData("Title", "404 Page Not Found!")
		ctx.View("shared/404.html")
	})

	// app.OnAnyErrorCode(func(ctx iris.Context) {
	// 	ctx.ViewData("Message", ctx.Values().
	// 		GetStringDefault("message", "The page you're looking for doesn't exist"))
	// 	ctx.View("shared/error.html")
	// })

	listen := fmt.Sprintf(":%d", 8080)
	log.Println("程序已经启动成功：您可以访问程序主页：http://localhost" + listen)
	app.Listen(listen)
}
