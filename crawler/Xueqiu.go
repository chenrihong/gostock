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

// Xueqiu 雪球的行情接口
type Xueqiu struct {
	baseURL      string
	ResponseData interface{}
	HomePageData []map[string]string
}

// 构造函数
func NewXueqiu() *Xueqiu {
	instance := new(Xueqiu)
	instance.baseURL = "https://stock.xueqiu.com"
	return instance
}

// GetIndexHq 指数行情
func (c *Xueqiu) GetIndexHq() error {
	return c.GetHq("SH000001,SZ399001,SZ399006,SH000688,HKHSI,HKHSCEI,HKHSCCI,.DJI,.IXIC,.INX")
}

// GetHq 发送请求
func (c *Xueqiu) GetHq(symbol string) error {

	rnd := strconv.FormatInt(time.Now().UnixNano(), 10)
	path := "/v5/stock/batch/quote.json?symbol=" + symbol + "&__t=" + rnd
	url := c.baseURL + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header = XueqiuHeader{}.HqHeader()
	req.Header.Set("path", path)

	xueqiucookie := XueqiuCookie{}
	cookies := xueqiucookie.Get()
	c.HomePageData = xueqiucookie.HomePageData

	for _, ck := range cookies {
		req.AddCookie(ck)
	}

	fmt.Println("正在请求雪球行情数据", url)

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

	jsonmap := map[string]interface{}{}
	json.Unmarshal(body, &jsonmap)
	c.ResponseData = jsonmap
	return nil
}
