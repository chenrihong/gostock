package controllers

import (
	"log"

	"github.com/chenrihong/gostock/controllers/funds"
	"github.com/chenrihong/gostock/controllers/stocks"
	"github.com/kataras/iris/v12"
)

// Routers 全局大模块路由注册
func RegRouters(app *iris.Application) {

	log.Println("注册程序路由 开始..")

	stocks.RegisteRouters(app)
	funds.RegisteRouters(app)

	log.Println("注册程序路由 完成..")
}
