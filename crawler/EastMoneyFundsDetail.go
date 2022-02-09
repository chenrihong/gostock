package crawler

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"gitee.com/chenrh/gotoolkit/conv"
	"github.com/PuerkitoBio/goquery"
	"github.com/chenrihong/gostock/models"
)

// 查询某基金前10持仓股，源网址
// http://fund.eastmoney.com/001298.html
type EastMoneyFundsDetails struct {
}

func (c *EastMoneyFundsDetails) GetData(fundCode string) ([]models.FUND_STOCK, error) {
	queryURL := fmt.Sprintf("http://fund.eastmoney.com/%s.html", fundCode)
	log.Println("爬取地址是：", queryURL)

	dom, err := goquery.NewDocument(queryURL)

	if err != nil {
		log.Println("爬取基金持仓数据时报错了：", err.Error())
		return nil, err
	}

	var list = []models.FUND_STOCK{}
	dom.Find("#position_shares tr").Each(func(i int, tr *goquery.Selection) {
		if i > 0 {
			one := models.FUND_STOCK{
				FUND_CODE: fundCode,
			}
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				if j == 0 {
					nodea := td.Find("a")
					href, has := nodea.Attr("href")
					if has {
						regs, _ := regexp.Compile(`\D`)
						code := regs.ReplaceAllString(href, "")
						one.STOCK_CODE = code
					}
					name, has := nodea.Attr("title")
					if has {
						one.STOCK_NAME = name
					}
				} else if j == 1 {
					ratestr := td.Text()
					ratestr = strings.Replace(ratestr, "%", "0", -1)
					var rate float64 = conv.ToFloat64(ratestr, 0)
					one.STOCK_RATE = rate
				}
			})

			if len(one.STOCK_CODE) != 0 && len(one.STOCK_NAME) != 0 {
				list = append(list, one)
			}
		}

	})
	return list, nil
}
