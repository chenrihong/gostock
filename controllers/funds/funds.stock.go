package funds

import (
	"fmt"
	"log"
	"time"

	"github.com/chenrihong/gostock/crawler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/xuri/excelize/v2"
)

// FundsStock serves the "/", "/ping" and "/hello".
type FundsStock struct {
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
func (c *FundsStock) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/stock/list", "StockList")
}

func (c *FundsStock) GetStock() mvc.Result {

	// 财报季数
	quarter := []string{"12-31", "09-30", "06-30", "03-31"}
	now := time.Now()

	list := []string{}

	year := now.Year()
	for i := 0; i < 3; i++ {
		for _, q := range quarter {
			quarterYear := year - i
			quarterString := fmt.Sprintf("%d-%s", quarterYear, q)
			dateQuarter, _ := time.Parse("2006-01-02", quarterString)

			if time.Now().After(dateQuarter) {
				list = append(list, quarterString)
			}
		}
	}
	log.Println(list)

	return mvc.View{
		Name: "funds/funds.stock.html",
		Data: iris.Map{
			"Title":   "基金持仓列表",
			"Quarter": list,
		},
	}
}

func (c *FundsStock) StockList() {

	condition := crawler.EastMoneyFundsStockQueryModel{}
	c.Ctx.ReadJSON(&condition)

	if condition.Export == 1 {
		c.StockListExport(&condition)
		return
	}

	craw := crawler.EastMoneyFundsStock{}
	body, _ := craw.GetQueryData(&condition)
	jsonmap, _ := craw.JsObjectToJSON(body)

	if jsonmap["data"] == nil {
		c.Ctx.JSON(iris.Map{"flag": true, "msg": "OK", "rows": nil, "total": 0, "jsonmap": jsonmap})
		return
	}

	rows := jsonmap["data"].([]interface{})
	pages := jsonmap["pages"]
	total := int(pages.(float64)) * condition.PageSize

	c.Ctx.JSON(iris.Map{"flag": true, "msg": "OK", "rows": rows, "total": total})
	return
}

// StockListExport  导出Excel文件
func (c *FundsStock) StockListExport(condition *crawler.EastMoneyFundsStockQueryModel) {
	condition.PageNo = 1
	condition.PageSize = 10000

	craw := crawler.EastMoneyFundsStock{}
	body, _ := craw.GetQueryData(condition)
	jsonmap, _ := craw.JsObjectToJSON(body)
	rows := jsonmap["data"].([]interface{})

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	var keys = []string{}
	var titles = []string{}

	keys = append(keys, "SECURITY_CODE")
	titles = append(titles, "股票代码")

	keys = append(keys, "SECURITY_NAME_ABBR")
	titles = append(titles, "股票名称")

	keys = append(keys, "REPORT_DATE")
	titles = append(titles, "财报季度")

	keys = append(keys, "ORG_TYPE_NAME")
	titles = append(titles, "机构类型")

	keys = append(keys, "HOULD_NUM")
	titles = append(titles, "机构家数")

	keys = append(keys, "HOLDCHA")
	titles = append(titles, "财季增减")

	keys = append(keys, "TOTAL_SHARES")
	titles = append(titles, "持股数量（股）")

	keys = append(keys, "HOLD_VALUE")
	titles = append(titles, "持股金额（元）")

	keys = append(keys, "HOLDCHA_NUM")
	titles = append(titles, "本季变动（股）")

	keys = append(keys, "HOLDCHA_RATIO")
	titles = append(titles, "变动比例(%)")

	// 写列头
	f.SetSheetRow("Sheet1", "A1", &titles)

	for rowIndex, row := range rows {
		// &[]interface{}{"1", nil, 2})
		var sheetData = []interface{}{}
		var mapdata = row.(map[string]interface{})

		for _, key := range keys {
			sheetData = append(sheetData, mapdata[key])
		}
		axis := fmt.Sprintf("A%d", rowIndex+2)
		f.SetSheetRow("Sheet1", axis, &sheetData)
	}

	// Save spreadsheet by the given path.
	path := "./static/logs/stockbook1.xlsx"
	if err := f.SaveAs(path); err != nil {
		fmt.Println(err)
	}

	c.Ctx.JSON(iris.Map{"flag": true, "msg": "OK", "rows": nil, "total": 0})
	return
}
