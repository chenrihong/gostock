package funds

import (
	"github.com/chenrihong/gostock/crawler"
	"github.com/chenrihong/gostock/ui"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// FundsList serves the "/", "/ping" and "/hello".
type FundsList struct {
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
func (c *FundsList) BeforeActivation(b mvc.BeforeActivation) {
	// b.Handle("POST", "/selfselect/add", "AddSelfSelectStock")
	// b.Handle("POST", "/crawler/fund/detail", "CrawlerFundDetails")
	b.Handle("GET", "/detail/{fundCode}", "FundDetails")
	b.Handle("POST", "/detail/{fundCode}/list", "FundDetailsList")
}

func (c *FundsList) GetList() mvc.Result {
	return mvc.View{
		Name: "funds/funds.list.html",
		Data: iris.Map{
			"Title": "基金列表",
		},
	}
}

// CrawlerFundStock 抓取基金持仓
func (c *FundsList) FundDetails() mvc.Result {
	code := c.Ctx.Params().Get("fundCode")

	// eastm := crawler.EastMoneyFundsDetails{}
	// list, _ = eastm.GetData(code)

	return mvc.View{
		Name: "funds/funds.detail.html",
		Data: iris.Map{
			"Title": code + " | 基金持仓情况",
			"Code":  code,
		},
	}
}

// PostList 基金列表数据源
func (c *FundsList) PostList() {

	condition := ui.BootstrapTablePage{}
	c.Ctx.ReadJSON(&condition)

	craw := crawler.EastMoneyFunds{}
	body, _ := craw.GetQueryData(&condition)
	jsonmap, _ := craw.JsObjectToJSON(body)

	rows := jsonmap["datas"].([]interface{})

	for i, row := range rows {
		rows[i] = iris.Map{"NAME": row}
	}

	total := jsonmap["allRecords"]

	c.Ctx.JSON(iris.Map{"flag": true, "msg": "OK", "rows": rows, "total": total})
	return
}

// FundDetailsList 基金持仓前10列表数据源
func (c *FundsList) FundDetailsList() {
	// code := c.Ctx.Params().Get("fundCode")

	// eastm := crawler.EastMoneyFundsDetails{}
	// list, err := eastm.GetData(code)

	// var total float64 = 0
	// for _, row := range list {
	// 	total += row.STOCK_RATE
	// }

	// list = append(list, models.FUND_STOCK{
	// 	ID:         -1,
	// 	FUND_CODE:  code,
	// 	STOCK_CODE: "TOTAL",
	// 	STOCK_NAME: "合计持仓",
	// 	STOCK_RATE: total,
	// })

	// if err != nil {
	// 	c.Ctx.JSON(iris.Map{"flag": false, "msg": err.Error(), "rows": nil, "total": 0})
	// 	return
	// }
	// c.Ctx.JSON(iris.Map{"flag": true, "msg": "OK", "rows": list, "total": len(list)})
	// return
}

// PostCrawler 抓取
func (c *FundsList) PostCrawler() {

	// sql := "select count(*) n from fund_info "
	// arr, _ := app.DB.Select(sql)

	// if len(arr) != 0 && arr[0]["n"].(int64) != 0 {
	// 	c.Ctx.JSON(iris.Map{"flag": true, "msg": fmt.Sprintf("之前已爬%d条，已经爬完，不必重新爬", arr[0]["n"])})
	// 	return
	// }

	// eastm := crawler.EastMoney{}
	// eastm.CrawlerFunds()

	// arr, _ = app.DB.Select(sql)

	// c.Ctx.JSON(iris.Map{"flag": true, "msg": fmt.Sprintf("爬取完成，共爬%d条", arr[0]["n"])})
	// return
}

// CrawlerFundStock 抓取基金持仓
func (c *FundsList) CrawlerFundDetails() {
	// sql := "select count(*) n from fund_info "
	// arr, _ := app.DB.Select(sql)

	// if len(arr) != 0 && arr[0]["n"].(int64) == 0 {
	// 	c.Ctx.JSON(iris.Map{"flag": false, "msg": "请先爬取基金数据"})
	// 	return
	// }

	// sql = "SELECT fund_code FROM fund_info"
	// arr, _ = app.DB.Select(sql)

	// fundCodes := []string{}
	// for _, m := range arr {
	// 	fundCodes = append(fundCodes, m["fund_code"].(string))
	// }

	// eastm := crawler.EastMoney{}
	// eastm.CrawlerFundDetail(fundCodes)

	// c.Ctx.JSON(iris.Map{"flag": true, "msg": "爬取完成"})
	// return
}
