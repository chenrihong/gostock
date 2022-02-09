package stocks

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/axgle/mahonia" // gbk to utf8
)

// CrawlerForStockSina 新浪的行情接口
type CrawlerForStockSina struct {
	SinaContent string
	Javascript  string
}

// SendRequest 发送请求
func (c *CrawlerForStockSina) SendRequest(url string) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	//提交请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	content := string(body)
	// 查看网页页面信息文字编码为gbk,而golang使用的utf8。
	// 加入下边代码将字符串编码有GBK转换为utf8
	utf8 := mahonia.NewDecoder("gbk").ConvertString(content)

	fmt.Println("SinaContent URL :" + url)
	fmt.Println("SinaContent :" + utf8)

	utf8 = strings.Replace(utf8, "\n", "", -1)
	c.SinaContent = utf8
	return nil
}
