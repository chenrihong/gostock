package stocks

import (
	"github.com/chenrihong/gostock/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// RegisteRouters 注册路由
func RegisteRouters(app *iris.Application) {

	p := mvc.New(app.Party("/stocks"))
	p.Router.Use(middleware.StocksNavMiddleware)

	p.Handle((new(StockIndex)))
	p.Handle(new(TodayController))
	p.Handle(new(PriceController))

	p.Handle(new(StockStrategyController))
	p.Handle(new(StockNoticeController))
	p.Handle(new(StockSector))
}
