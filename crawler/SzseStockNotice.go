package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type SzseStockNoticeSearchModel struct {
	ChannelCode []string `json:"channelCode"`
	Stock       []string `json:"stock"`
	SearchKey   []string `json:"searchKey"`
	PageNum     int      `json:"pageNum"`
	PageSize    int      `json:"pageSize"`
	SeDate      []string `json:"seDate"`
}

// SzseStockNotice 爬取公司公告
type SzseStockNotice struct {
	Condition     *StockNoticeConditionModel
	postCondition *SzseStockNoticeSearchModel
}

func (c *SzseStockNotice) GetURLQueryParams() string {
	condition := c.Condition

	c.postCondition = &SzseStockNoticeSearchModel{}

	if len(condition.START_DATE) == 0 {
		condition.START_DATE = time.Now().Add(-2 * 24 * time.Hour).Format("2006-01-02")
	}
	if len(condition.END_DATE) == 0 {
		condition.END_DATE = time.Now().Format("2006-01-02")
	}

	c.postCondition.ChannelCode = []string{"listedNotice_disc"}
	c.postCondition.SeDate = []string{condition.START_DATE, condition.END_DATE}
	c.postCondition.PageNum = condition.PageNo
	c.postCondition.PageSize = condition.PageSize

	if len(condition.Search) != 0 {
		c.postCondition.SearchKey = []string{condition.Search}
	}
	if len(condition.CODE) != 0 {
		c.postCondition.Stock = []string{condition.CODE}
	}

	stockListURL := `http://www.szse.cn/api/disc/announcement/annList`
	stockListURL += fmt.Sprintf("&random=0.%d", time.Now().UnixNano())

	log.Println("爬取地址是：", stockListURL)
	return stockListURL
}

func (c *SzseStockNotice) GetQueryJSON() map[string]interface{} {
	data, _ := c.GetQueryData()
	m := map[string]interface{}{}
	json.Unmarshal(data, &m)
	return m
}

// GetQueryData 发送请求
func (c *SzseStockNotice) GetQueryData() ([]byte, error) {

	url := c.GetURLQueryParams()

	postjson, err := json.Marshal(c.postCondition)
	if err != nil {
		return nil, err
	}

	content := string(postjson)

	fmt.Println(content)

	req, err := http.NewRequest("POST", url, strings.NewReader(content))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "http://www.szse.cn/disclosure/listed/notice/index.html")
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

	// log.Println(string(body))
	return body, nil
}
