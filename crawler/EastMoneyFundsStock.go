// http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=gp&rs=&gs=0&sc=jnzf&st=desc&sd=2020-07-23&ed=2021-07-23&qdii=&tabSubtype=,,,,,&pi=3&pn=50&dx=1&v=0.6518763607107061
package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chenrihong/gostock/ui"
)

// 查询基金持仓排名靠前的股票
// http://data.eastmoney.com/zlsj
// http://data.eastmoney.com/zlsj/2021-06-30-1-2.html

type EastMoneyFundsStockQueryModel struct {
	ui.BootstrapTablePage

	Code       string
	Type       string // 1 基金持仓， 2 QFII 持仓， 3社保持仓
	Zjc        string
	ReportDate string
	Export     int // 1 导出操作 0 查询
}

type EastMoneyFundsStock struct {
}

// 源网址
// http://data.eastmoney.com/dataapi/zlsj/list?
// tkn=eastmoney&date=2021-06-30
// &code=&type=1&zjc=1&sortField=HOLD_VALUE&sortDirec=1&pageNum=1&pageSize=50&cfg=jjsjtj&p=1&pageNo=1

func (c *EastMoneyFundsStock) GetURLQueryParams(condition *EastMoneyFundsStockQueryModel) string {

	stockListURL := "http://data.eastmoney.com/dataapi/zlsj/list?"

	stockListURL += fmt.Sprintf("&date=%s", condition.ReportDate)
	stockListURL += fmt.Sprintf("&code=%s", condition.Code)
	stockListURL += fmt.Sprintf("&type=%s", condition.Type)
	stockListURL += fmt.Sprintf("&zjc=%s", condition.Zjc)

	if len(condition.Sort) > 0 {
		stockListURL += fmt.Sprintf("&sortField=%s", condition.Sort)
	}

	order := 1
	if condition.Order == "desc" {
		order = 1
	} else if condition.Order == "asc" {
		order = 0
	}

	stockListURL += fmt.Sprintf("&sortDirec=%d", order)

	stockListURL += fmt.Sprintf("&pageNum=%d", condition.PageNo)
	stockListURL += fmt.Sprintf("&pageSize=%d", condition.PageSize)

	stockListURL += "&cfg=jjsjtj&p=1"
	stockListURL += fmt.Sprintf("&pageNo=%d", condition.PageNo)

	log.Println("爬取地址是：", stockListURL)
	return stockListURL
}

// GetData 发送请求
func (c *EastMoneyFundsStock) GetQueryData(condition *EastMoneyFundsStockQueryModel) ([]byte, error) {

	url := c.GetURLQueryParams(condition)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "http://data.eastmoney.com/zlsj/2021-06-30-1-2.html")
	req.Header.Set("Host", "data.eastmoney.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("StatusCode:%d Status:%s", resp.StatusCode, resp.Status))
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// JsObjectToJSON 流转化为对象
func (c *EastMoneyFundsStock) JsObjectToJSON(body []byte) (map[string]interface{}, error) {
	jsonmap := map[string]interface{}{}
	err := json.Unmarshal(body, &jsonmap)

	if err != nil {
		return nil, err
	}
	return jsonmap, nil
}
