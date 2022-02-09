package stocks

import (
	"github.com/chenrihong/gostock/crawler"
	"github.com/chenrihong/gostock/ui"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// StockSector serves the "/", "/ping" and "/hello".
type StockSector struct {
	// Optionally: context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	//Service services.UserService
}

// BeforeActivation 勾子
func (c *StockSector) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/hot_sector", "GetHotSectorToday")
}

// Get handles GET: http://localhost:8080/stocks/today.
func (c *StockSector) GetSector() mvc.Result {
	return mvc.View{
		Name: "stocks/stocks.sector.html",
		Data: iris.Map{
			"Title": "板块行情",
		},
	}
}

// GetHotSectorToday 今日热门板块
func (c *StockSector) GetHotSectorToday() {
	condition := ui.BootstrapTablePage{}
	c.Ctx.ReadJSON(&condition)

	craw := crawler.EastMoneyStockSector{}
	body, _ := craw.GetQueryData(&condition)
	jsonmap, _ := craw.JsObjectToJSON(body)

	data := jsonmap["data"].(map[string]interface{})

	rows := data["diff"].([]interface{})
	total := data["total"]
	c.Ctx.JSON(iris.Map{"flag": true, "msg": "OK", "rows": rows, "total": total})
	return
}
