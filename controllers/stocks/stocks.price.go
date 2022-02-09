package stocks

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"fmt"
	"strings"
	"time"
)

// PriceController serves the "/", "/ping" and "/hello".
type PriceController struct {
	// Optionally: context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	//Service services.UserService
}

// GetPrice handles GET: http://localhost:8080/example/.
func (c *PriceController) GetPrice() mvc.Result {

	lists := []string{"sz002230", "sz000002", "sz000001"}
	opcodes := []string{}
	jsobj := "var opcodeobj = {};"
	for _, item := range lists {
		code := item
		opcodes = append(opcodes, code)
		jsobj += "opcodeobj." + code + " = hq_str_" + code + ";"
	}

	return mvc.View{
		Name: "stocks/stocks.price.html",
		Data: iris.Map{
			"Title":       "股票计价器",
			"OPCODE":      strings.Join(opcodes, ","),
			"LIST_OPCODE": opcodes,
			"opcodeobj":   jsobj,
		},
	}
}

// isMarketRestNow 当前时间是否正在开市
func (c *PriceController) isMarketRestNow() bool {
	// 如果不是开市时间，不进行轮询
	now := time.Now().Local()
	week := int(now.Weekday())
	if week == 0 || week == 6 {
		return true
	}

	format := "2006-1-2 15:04:05"
	var todaystr = now.Format("2006-01-02")

	before915 := fmt.Sprintf("%s 09:14:59", todaystr)
	// before915dt, _ := time.Parse(format, before915) // 有时区问题
	before915dt, _ := time.ParseInLocation(format, before915, time.Local)

	after1130 := fmt.Sprintf("%s 11:30:00", todaystr)
	after1130dt, _ := time.ParseInLocation(format, after1130, time.Local)

	before1300 := fmt.Sprintf("%s 13:00:00", todaystr)
	before1300dt, _ := time.ParseInLocation(format, before1300, time.Local)

	after1500 := fmt.Sprintf("%s 15:00:00", todaystr)
	after1500dt, _ := time.ParseInLocation(format, after1500, time.Local)

	if now.Before(before915dt) {
		return true
	}
	if now.After(after1130dt) && now.Before(before1300dt) {
		return true
	}
	if now.After(after1500dt) {
		return true
	}
	return false
}

// GetHq http://localhost:8080/stocks/hq?opcodes=s_sh000001,s_sz399001,s_sz399006,sz002142,sz002602
func (c *PriceController) GetHq() {
	// 自选股
	opcodes := c.Ctx.URLParam("opcodes")

	hqurl := "http://hq.sinajs.cn/list=" + opcodes

	cls := CrawlerForStockSina{}
	err := cls.SendRequest(hqurl)
	if err != nil {
		c.Ctx.JSON(map[string]interface{}{"flag": false, "msg": err, "data": "", "is_open": false})
		return
	}

	var content = cls.SinaContent
	var isOpengo = c.isMarketRestNow() == false
	json := map[string]interface{}{"flag": true, "msg": "", "data": content, "is_open": isOpengo}
	c.Ctx.JSON(json)
	return
}
