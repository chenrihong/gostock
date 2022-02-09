package stocks

import (
	"strings"

	"github.com/chenrihong/gostock/crawler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// TodayController serves the "/", "/ping" and "/hello".
type TodayController struct {
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
func (c *TodayController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/hot_today", "GetHotToday")
}

// Get handles GET: http://localhost:8080/stocks/today.
func (c *TodayController) GetToday() mvc.Result {
	return mvc.View{
		Name: "stocks/stocks.today.html",
		Data: iris.Map{
			"Title": "今日行情",
		},
	}
}

// GetHotToday 今日热门股
func (c *TodayController) GetHotToday() {
	cls := crawler.XueqiuStockRanking{}
	tabid := c.Ctx.FormValue("tabid")
	size := c.Ctx.FormValue("size")
	arr := strings.Split(tabid, "_")
	typestr := "sha"
	order := "desc"

	if len(arr) == 2 {
		typestr = arr[0]
		order = arr[1]
	}

	err := cls.SendRequest(typestr, size, order)
	if err != nil {
		c.Ctx.JSON(iris.Map{"flag": false, "msg": err})
		return
	}
	c.Ctx.JSON(iris.Map{"flag": true, "msg": "", "data": cls.ResponseData})
	return
}
