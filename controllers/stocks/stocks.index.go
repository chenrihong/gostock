package stocks

import (
	"github.com/chenrihong/gostock/crawler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// StockIndex serves the "/", "/ping" and "/hello".
type StockIndex struct {
	// Optionally: context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	//Service services.UserService
}

// BeforeActivation Mvc 注册路由
func (c *StockIndex) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/indexhq", "GetIndexToday")
}

// Get handles GET: http://localhost:8080/stocks.
func (c *StockIndex) Get() mvc.Result {
	return mvc.View{
		Name: "stocks/stocks.index.html",
		Data: iris.Map{
			"Title": "证券分析首页",
		},
	}
}

// GetIndexToday 今日指数指标
func (c *StockIndex) GetIndexToday() {
	cls := crawler.NewXueqiu()

	err := cls.GetIndexHq()

	if err != nil {
		c.Ctx.JSON(iris.Map{"flag": false, "msg": err})
		return
	}

	c.Ctx.JSON(iris.Map{"flag": true, "msg": "", "data": cls.ResponseData, "xueqiu": cls.HomePageData})
	return
}
