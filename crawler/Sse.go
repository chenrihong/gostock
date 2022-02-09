package crawler

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"

	"github.com/chenrihong/gostock/models"
)

// 上交所爬虫
type Sse struct {
	/*
		WaitGroup：同步等待组
		可以使用Add(),设置等待组中要 执行的子goroutine的数量，
		使用wait(),让主程序处于等待状态。直到等待组中子程序执行完毕。解除阻塞		​
		子gorotuine对应的函数中。wg.Done()，用于让等待组中的子程序的数量减1
	*/
	wg sync.WaitGroup
}

// CrawlerStocks 爬取上交所股票列表
func (c *Sse) CrawlerStocks() {

	ssesl := SseStockList{}
	count := ssesl.GetStockCount()
	var pagesize = 100
	var pagecount = int(math.Ceil(float64(count) / float64(pagesize)))

	c.wg.Add(pagecount)

	for i := 0; i < pagecount; i++ {
		go func(index int) {
			fmt.Println("正在拉取第", index, "页，一共", pagecount, "页")
			data, err := ssesl.GetQueryData(index, pagesize)
			if err == nil {
				m := map[string]interface{}{}
				json.Unmarshal(data, &m)

				pageHelp := m["pageHelp"].(map[string]interface{})
				result := pageHelp["data"].([]interface{})

				stockModels := []models.STOCK_INFO{}
				for _, item := range result {
					mapobj := item.(map[string]interface{})
					var model models.STOCK_INFO
					model.COMPANY_CODE = mapobj["COMPANY_CODE"].(string)
					model.COMPANY_ABBR = mapobj["COMPANY_ABBR"].(string)
					model.LISTING_DATE = mapobj["LISTING_DATE"].(string)
					model.COMPANY_KIND = "sh"
					stockModels = append(stockModels, model)
				}
			}

			c.wg.Done()
		}(i + 1)
	}

	c.wg.Wait()

	fmt.Println("上交所数据已拉取完毕喽")
}
