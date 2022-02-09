package stocks

import (
	"github.com/chenrihong/gostock/crawler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// StockNoticeController serves the "/", "/ping" and "/hello".
type StockNoticeController struct {
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
func (c *StockNoticeController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Handle("POST", "/selfselect/add", "AddSelfSelectStock")
	// b.Handle("POST", "/selfselect/remove", "RemoveSelfSelectStock")
}

func (c *StockNoticeController) GetNotice() mvc.Result {
	return mvc.View{
		Name: "stocks/stocks.notice.html",
		Data: iris.Map{
			"Title": "股票公告列表",
		},
	}
}

func (c *StockNoticeController) PostNotice() {
	condition := crawler.StockNoticeConditionModel{}
	c.Ctx.ReadJSON(&condition)

	if condition.KIND == "sh" {
		crawlerNotice := crawler.SseStockNotice{
			Condition: &condition,
		}

		jsonrets := crawlerNotice.GetQueryJSON()

		rows := jsonrets["result"].([]interface{})
		pageHelp := jsonrets["pageHelp"].(map[string]interface{})
		total := pageHelp["total"]

		c.Ctx.JSON(iris.Map{"flag": true, "msg": "OK", "rows": rows, "total": total})
		return
	}

	if condition.KIND == "sz" {
		crawlerNotice := crawler.SzseStockNotice{
			Condition: &condition,
		}

		jsonrets := crawlerNotice.GetQueryJSON()
		rows := jsonrets["data"].([]interface{})
		total := jsonrets["announceCount"]

		c.Ctx.JSON(iris.Map{"flag": true, "msg": "OK", "rows": rows, "total": total})
		return
	}
}

// PostCrawler 抓取
// func (c *StockNoticeController) PostCrawler() {

// 	sse := crawler.Sse{}
// 	sse.CrawlerStocks()

// 	szse := crawler.Szse{}
// 	szse.CrawlerStocks()

// 	c.Ctx.JSON(iris.Map{"flag": true, "msg": ""})
// 	return
// }
