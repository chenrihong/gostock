// http://data.eastmoney.com/bkzj/hy.html
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

// 查询当日板块行业行情

type EastMoneyStockSector struct {
}

// 源网址
// http://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=50&pn=2&np=1&fltt=2&invt=2&ut=b2884a393a59ad64002292a3e90d46a5
//&fs=m%3A90+t%3A2&fields=f12%2Cf14%2Cf2%2Cf3%2Cf62%2Cf184%2Cf66%2Cf69%2Cf72%2Cf75%2Cf78%2Cf81%2Cf84%2Cf87%2Cf204%2Cf205%2Cf124%2Cf1%2Cf13
func (c *EastMoneyStockSector) GetURLQueryParams(condition *ui.BootstrapTablePage) string {

	stockListURL := "http://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1"

	stockListURL += fmt.Sprintf("&pz=%d", condition.PageSize)
	stockListURL += fmt.Sprintf("&pn=%d", condition.PageNo)

	stockListURL += "&np=1&fltt=2&invt=2&ut=b2884a393a59ad64002292a3e90d46a5"
	// stockListURL += "&fs=m:90 t:2" // 这里要转码，不然报错 502
	stockListURL += "&fs=m%3A90+t%3A2"

	stockListURL += "&fields=f12,f14,f2,f3,f62,f184,f66,f69,f72,f75,f78,f81,f84,f87,f204,f205,f124,f1,f13"

	log.Println("爬取地址是：", stockListURL)
	return stockListURL
}

// GetData 发送请求
func (c *EastMoneyStockSector) GetQueryData(condition *ui.BootstrapTablePage) ([]byte, error) {

	url := c.GetURLQueryParams(condition)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "http://data.eastmoney.com/")
	req.Header.Set("Host", "push2.eastmoney.com")
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
func (c *EastMoneyStockSector) JsObjectToJSON(body []byte) (map[string]interface{}, error) {
	jsonmap := map[string]interface{}{}
	err := json.Unmarshal(body, &jsonmap)

	if err != nil {
		return nil, err
	}
	return jsonmap, nil
}
