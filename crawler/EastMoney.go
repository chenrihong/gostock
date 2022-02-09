package crawler

import (
	"fmt"
	"sync"
)

// 东方财富，天天基金网，爬虫
type EastMoney struct {
	/*
		WaitGroup：同步等待组
		可以使用Add(),设置等待组中要 执行的子goroutine的数量，
		使用wait(),让主程序处于等待状态。直到等待组中子程序执行完毕。解除阻塞		​
		子gorotuine对应的函数中。wg.Done()，用于让等待组中的子程序的数量减1
	*/
	wg sync.WaitGroup
}

// CrawlerFunds 爬取天天基金网基金列表
func (c *EastMoney) CrawlerFunds() {

	// fundsCrawler := EastMoneyFunds{}
	// count := fundsCrawler.GetRecordsCount()
	// var pagesize = 100
	// var pagecount = int(math.Ceil(float64(count) / float64(pagesize)))

	// c.wg.Add(pagecount)

	// db := app.GetDB()

	// // 清空原有数据
	// db.Exec("TRUNCATE TABLE fund_info")

	// now := time.Now()

	// for i := 0; i < pagecount; i++ {
	// 	go func(index int) {
	// 		fmt.Println("正在拉取第", index, "页，共", pagecount, "页")
	// 		condition := ui.BootstrapTablePage{
	// 			PageNo:   index,
	// 			PageSize: pagesize,
	// 		}
	// 		data, err := fundsCrawler.GetQueryData(&condition)
	// 		if err == nil {
	// 			jsonmap, err := fundsCrawler.JsObjectToJSON(data)

	// 			if err == nil {
	// 				datas := jsonmap["datas"].([]interface{})

	// 				fund := []models.FUND_INFO{}
	// 				for _, item := range datas {

	// 					//"004616, 0
	// 					//中欧电子信息产业沪港深股票A, 1
	// 					//ZODZXXCYHGSGPA,2
	// 					//2021-07-23,3
	// 					//3.0390,4
	// 					//3.0390,5
	// 					//1.28, 6 日增长率
	// 					//9.97, 7近一周
	// 					//10.55,8近一月
	// 					//35.53,9 近三月
	// 					//22.86, 10 近6月
	// 					// 41.98, 11 近一年
	// 					// 158.07, 12 近2年
	// 					//186.89, 13 近3年
	// 					//35.82,14 今年以来
	// 					//203.90,15 成立以来
	// 					//2017-07-07,1,50.0815,1.50%,0.15%,1,0.15%,1,"

	// 					arr := strings.Split(item.(string), ",")

	// 					model := models.FUND_INFO{
	// 						FUND_CODE:      arr[0],
	// 						FUND_NAME:      arr[1],
	// 						FUND_DWJZ:      conv.ToFloat64(arr[4], 0),
	// 						FUND_LJJZ:      conv.ToFloat64(arr[5], 0),
	// 						RATE_1WEEK:     conv.ToFloat64(arr[7], 0),
	// 						RATE_1MONTH:    conv.ToFloat64(arr[8], 0),
	// 						RATE_3MONTH:    conv.ToFloat64(arr[9], 0),
	// 						RATE_6MONTH:    conv.ToFloat64(arr[10], 0),
	// 						RATE_1YEAR:     conv.ToFloat64(arr[11], 0),
	// 						RATE_2YEAR:     conv.ToFloat64(arr[12], 0),
	// 						RATE_3YEAR:     conv.ToFloat64(arr[13], 0),
	// 						RATE_THIS_YEAR: conv.ToFloat64(arr[14], 0),
	// 						RATE_ALL_TIME:  conv.ToFloat64(arr[15], 0),
	// 						CREATED_ON:     &now,
	// 					}

	// 					fund = append(fund, model)
	// 				}
	// 				db.Create(&fund)
	// 			}

	// 		}

	// 		c.wg.Done()
	// 	}(i + 1)
	// }

	// c.wg.Wait()

	fmt.Println("天天基金网基金列表数据已拉取完毕喽")
}

// CrawlerFunds 爬取天天基金网基金列表
func (c *EastMoney) CrawlerFundDetail(fundCodes []string) {

	fundsCrawler := EastMoneyFundsDetails{}
	pagecount := len(fundCodes)
	c.wg.Add(pagecount)

	for i := 0; i < pagecount; i++ {
		fundCode := fundCodes[i]

		go func(index int) {
			fmt.Println("正在拉取第", index, "只，共", pagecount, "只")
			fundsCrawler.GetData(fundCode)
			c.wg.Done()
		}(i + 1)
	}

	c.wg.Wait()

	fmt.Println("天天基金网基金列表数据已拉取完毕喽")
}
