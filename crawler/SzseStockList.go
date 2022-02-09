package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SzseStockList struct {
}

// 深交所 PageSize 参数没有，固定的 20。
func (c *SzseStockList) GetURLQueryParams(pageNO int) string {
	stockListURL := "http://www.szse.cn/api/report/ShowReport/data?SHOWTYPE=JSON&CATALOGID=1110&TABKEY=tab1"
	stockListURL += fmt.Sprintf("&PAGENO=%d", pageNO)
	stockListURL += fmt.Sprintf("&random=0.%d", time.Now().UnixNano())
	log.Println("爬取地址是：", stockListURL)
	return stockListURL
}

// GetPageCount 查下数据量，以便分页精准拉取
func (c *SzseStockList) GetPageCount() int {
	body, err := c.GetQueryData(1)
	if err != nil {
		return 0
	}
	szjson := []map[string]interface{}{}
	json.Unmarshal(body, &szjson)
	ele0 := szjson[0]
	metadata := ele0["metadata"].(map[string]interface{})
	pagecount := metadata["pagecount"].(float64)
	return int(pagecount)
}

// GetData 发送请求
func (c *SzseStockList) GetQueryData(pageNO int) ([]byte, error) {

	url := c.GetURLQueryParams(pageNO)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "http://www.szse.cn/market/product/stock/list/index.html")
	req.Header.Set("Host", "www.szse.cn")
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
