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

// SseStockList 爬取公司列表
type SseStockList struct {
}

func (c *SseStockList) GetURLQueryParams(pageNO, pageSize int) string {
	stockListURL := "http://query.sse.com.cn/security/stock/getStockListData2.do?isPagination=true&stockCode=&csrcCode=&areaName=&stockType=1&pageHelp.cacheSize=1"
	stockListURL += fmt.Sprintf("&pageHelp.beginPage=%d", pageNO)
	stockListURL += fmt.Sprintf("&pageHelp.pageNo=%d", pageNO)
	stockListURL += fmt.Sprintf("&pageHelp.pageSize=%d", pageSize)
	stockListURL += fmt.Sprintf("&pageHelp.endPage=%d1", pageNO)
	stockListURL += fmt.Sprintf("&_=%d", time.Now().UnixNano())

	log.Println("爬取地址是：", stockListURL)

	return stockListURL
}

// GetStockCount 查下数据量，以便分页精准拉取
func (c *SseStockList) GetStockCount() int {
	body, err := c.GetQueryData(1, 1)
	if err != nil {
		log.Println("爬取上交所股票数据时出错：", err.Error())
		return 0
	}

	datamap := map[string]interface{}{}
	json.Unmarshal(body, &datamap)

	pageHelp := datamap["pageHelp"].(map[string]interface{})
	result := pageHelp["total"].(float64)

	return int(result)
}

// GetData 发送请求
func (c *SseStockList) GetQueryData(pageNO, pageSize int) ([]byte, error) {

	url := c.GetURLQueryParams(pageNO, pageSize)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "http://www.sse.com.cn/")
	req.Header.Set("Host", "query.sse.com.cn")
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

	// log.Println(string(body))
	return body, nil
}
