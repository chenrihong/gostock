package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/chenrihong/gostock/ui"
)

// StockNoticeConditionModel 上市公司公告查询条件
type StockNoticeConditionModel struct {
	ui.BootstrapTablePage

	START_DATE string
	END_DATE   string
	CODE       string
	KIND       string
}

// SseStockNotice 爬取公司公告
type SseStockNotice struct {
	Condition *StockNoticeConditionModel
}

func (c *SseStockNotice) GetURLQueryParams() string {
	condition := c.Condition

	if len(condition.START_DATE) == 0 {
		condition.START_DATE = time.Now().Add(-2 * 24 * time.Hour).Format("2006-01-02")
	}
	if len(condition.END_DATE) == 0 {
		condition.END_DATE = time.Now().Format("2006-01-02")
	}

	stockListURL := `http://query.sse.com.cn/commonQuery.do?isPagination=true&pageHelp.cacheSize=1&type=inParams&sqlId=COMMON_PL_SSGSXX_ZXGG_L&BULLETIN_TYPE=`

	stockListURL += fmt.Sprintf("&pageHelp.pageNo=%d", condition.PageNo)
	stockListURL += fmt.Sprintf("&pageHelp.pageSize=%d", condition.PageSize)
	stockListURL += fmt.Sprintf("&pageHelp.beginPage=%d", condition.PageNo)
	stockListURL += fmt.Sprintf("&pageHelp.endPage=%d", condition.PageNo)

	stockListURL += fmt.Sprintf("&START_DATE=%s", condition.START_DATE)
	stockListURL += fmt.Sprintf("&END_DATE=%s", condition.END_DATE)
	stockListURL += fmt.Sprintf("&SECURITY_CODE=%s", condition.CODE)
	stockListURL += fmt.Sprintf("&TITLE=%s", condition.Search)
	stockListURL += fmt.Sprintf("&_=%d", time.Now().UnixNano())

	log.Println("爬取地址是：", stockListURL)
	return stockListURL
}

func (c *SseStockNotice) GetQueryJSON() map[string]interface{} {
	data, _ := c.GetQueryData()
	m := map[string]interface{}{}
	json.Unmarshal(data, &m)
	return m
}

// GetData 发送请求
func (c *SseStockNotice) GetQueryData() ([]byte, error) {

	url := c.GetURLQueryParams()

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
