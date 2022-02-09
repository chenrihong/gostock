package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// XueqiuStockRanking 雪球股票涨幅跌幅榜单
type XueqiuStockRanking struct {
	ResponseData interface{}
}

// SendRequest 发送请求
func (c *XueqiuStockRanking) SendRequest(typestr, size, order string) error {
	rnd := strconv.FormatInt(time.Now().UnixNano(), 10)

	url := "https://xueqiu.com/service/v5/stock/screener/quote/list?page=1&size=" + size
	url += "&order=" + order
	url += "&type=" + typestr
	url += "&order_by=percent&exchange=CN&market=CN&_=" + rnd

	fmt.Println("xueqiu.com hq:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("StatusCode:%d Status:%s", resp.StatusCode, resp.Status))
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	m := map[string]interface{}{}
	json.Unmarshal(body, &m)

	c.ResponseData = m
	return nil
}
