package crawler

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sync"

	"github.com/chenrihong/gostock/models"
)

// 深交所爬虫
type Szse struct {
	/*
		WaitGroup：同步等待组
		可以使用Add(),设置等待组中要 执行的子goroutine的数量，
		使用wait(),让主程序处于等待状态。直到等待组中子程序执行完毕。解除阻塞		​
		子gorotuine对应的函数中。wg.Done()，用于让等待组中的子程序的数量减1
	*/
	wg sync.WaitGroup
}

// CrawlerStocks 爬取深交所股票列表
func (c *Szse) CrawlerStocks() {

	szsesl := SzseStockList{}
	pagecount := szsesl.GetPageCount()

	c.wg.Add(pagecount)

	for i := 0; i < pagecount; i++ {
		go func(index int) {
			fmt.Println("正在拉取第", index, "页，一共", pagecount, "页")
			bytedata, err := szsesl.GetQueryData(index)
			if err == nil {
				szjson := []map[string]interface{}{}
				json.Unmarshal(bytedata, &szjson)
				ele0 := szjson[0]
				data := ele0["data"].([]interface{})

				stockModels := []models.STOCK_INFO{}
				for _, item := range data {
					mapobj := item.(map[string]interface{})

					html := mapobj["agjc"].(string)
					reg := regexp.MustCompile(`<u>(?s:(.*?))</u>`)
					if reg == nil {
						fmt.Println("正则表达式编译有误！正则存在问题：", "<u>(?s:(.*?))</u>")
						continue
					}
					result := reg.FindAllStringSubmatch(html, -1)
					//result =  [[<u>平安银行</u> 平安银行]]
					for _, text := range result {
						html = text[1]
						break
					}

					var model models.STOCK_INFO
					model.COMPANY_CODE = mapobj["agdm"].(string)
					model.COMPANY_ABBR = html
					model.COMPANY_KIND = "sz"

					if mapobj["sshymc"] != nil {
						model.COMPANY_ATTR = mapobj["sshymc"].(string)
					}
					if mapobj["agssrq"] != nil {
						model.LISTING_DATE = mapobj["agssrq"].(string)
					}
					stockModels = append(stockModels, model)
				}

			}

			c.wg.Done()
		}(i + 1)
	}

	c.wg.Wait()

	fmt.Println("深交所数据已拉取完毕喽")
}
