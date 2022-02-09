package stocks

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// StockStrategyController serves the "/", "/ping" and "/hello".
type StockStrategyController struct {
	// Optionally: context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding).
	Ctx iris.Context
	// Our UserService, it's an interface which
	// is binded from the main application.
	//Service services.UserService
}

// BeforeActivation 注册自定义的路由
func (c *StockStrategyController) BeforeActivation(b mvc.BeforeActivation) {

}

// GetStrategy  股票操作策略界面
func (c *StockStrategyController) GetStrategy() mvc.Result {
	return mvc.View{
		Name: "stocks/stocks.strategy.html",
		Data: iris.Map{
			"Title": "股票操作策略",
		},
	}
}
