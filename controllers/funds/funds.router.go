package funds

import (
	"github.com/chenrihong/gostock/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// RegisteRouters 注册路由
func RegisteRouters(app *iris.Application) {

	p := mvc.New(app.Party("/funds"))
	p.Router.Use(middleware.StocksNavMiddleware)

	p.Handle(new(FundsList))
	p.Handle(new(FundsStock))
}
