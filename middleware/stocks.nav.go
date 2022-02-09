package middleware

import (
	// "fmt"

	"html/template"

	"github.com/kataras/iris/v12"
)

// StocksNavMiddleware 中间件
func StocksNavMiddleware(ctx iris.Context) {
	html := `
	<ul class="nav justify-content-end mb-2">
		<li class="nav-item">
		<a class="nav-link" href="/stocks/today">今日行情</a>
		</li>
		<li class="nav-item">
		<a class="nav-link" href="/stocks/sector">今日板块</a>
		</li>
		<li class="nav-item">
		<a class="nav-link" href="/stocks/calendar">投资日历</a>
		</li>
		<li class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" href="#" id="navbarDarkDropdownMenuLink1" role="button" data-bs-toggle="dropdown" aria-expanded="false">
            股票
          </a>
          <ul class="dropdown-menu dropdown-menu-light" aria-labelledby="navbarDarkDropdownMenuLink1">
            <li><a class="dropdown-item" href="/stocks">行情概览</a></li>
            <li><a class="dropdown-item" href="/stocks/list">股票列表</a></li>
            <li><a class="dropdown-item" href="/stocks/notice">股票公告</a></li>
            <li><a class="dropdown-item" href="/stocks/weigh/index">指数分析</a></li>
          </ul>
        </li>
		<li class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" href="#" id="navbarDarkDropdownMenuLink" role="button" data-bs-toggle="dropdown" aria-expanded="false">
            基金
		  </a>
          <ul class="dropdown-menu dropdown-menu-light" aria-labelledby="navbarDarkDropdownMenuLink">
            <li><a class="dropdown-item" href="/funds/list">基金排名</a></li>
            <li><a class="dropdown-item" href="/funds/stock">机构持股</a></li>
          </ul>
        </li>
		<li class="nav-item">
		<a class="nav-link" href="/stocks/strategy">操作策略</a>
		</li>
	</ul>`

	ctx.ViewData("StocksModuleNav", template.HTML(html))

	ctx.Next()
}
