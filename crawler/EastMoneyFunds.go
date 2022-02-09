// http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=gp&rs=&gs=0&sc=jnzf&st=desc&sd=2020-07-23&ed=2021-07-23&qdii=&tabSubtype=,,,,,&pi=3&pn=50&dx=1&v=0.6518763607107061
package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/chenrihong/gostock/ui"
)

type EastMoneyFunds struct {
}

// 源网址
// http://fund.eastmoney.com/data/fundranking.html#tgp;c0;r;sjnzf;pn50;ddesc;qsd20200723;qed20210723;qdii;zq;gg;gzbd;gzfs;bbzt;sfbb

func (c *EastMoneyFunds) GetURLQueryParams(condition *ui.BootstrapTablePage) string {

	today := time.Now().Format("2006-01-02")

	if len(condition.Sort) == 0 {
		condition.Sort = "zzf" // 周涨幅
	}
	if len(condition.Order) == 0 {
		condition.Order = "desc"
	}

	stockListURL := "http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=gp&rs=&gs=0&qdii=&tabSubtype=,,,,,&dx=1"
	stockListURL += fmt.Sprintf("&sc=%s", condition.Sort)
	stockListURL += fmt.Sprintf("&st=%s", condition.Order)
	stockListURL += fmt.Sprintf("&sd=%s", today)
	stockListURL += fmt.Sprintf("&ed=%s", today)
	stockListURL += fmt.Sprintf("&pi=%d", condition.PageNo)
	stockListURL += fmt.Sprintf("&pn=%d", condition.PageSize)
	stockListURL += fmt.Sprintf("&v=0.%d", time.Now().UnixNano())
	log.Println("爬取地址是：", stockListURL)
	return stockListURL
}

// GetPageCount 查下数据量，以便分页精准拉取
func (c *EastMoneyFunds) GetRecordsCount() int {
	first := ui.BootstrapTablePage{
		PageNo:   1,
		PageSize: 1,
	}
	body, err := c.GetQueryData(&first)
	if err != nil {
		// app.Logger().Errorf("爬取天天基金数据时出错：", err.Error())
		return 0
	}

	fundsjson, err := c.JsObjectToJSON(body)
	if err != nil {
		// app.Logger().Errorf("处理天天基金数据出错", err.Error())
		return 0
	}

	count := fundsjson["allRecords"].(float64)
	return int(count)
}

// GetData 发送请求
func (c *EastMoneyFunds) GetQueryData(condition *ui.BootstrapTablePage) ([]byte, error) {

	url := c.GetURLQueryParams(condition)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "http://fund.eastmoney.com/data/fundranking.html")
	req.Header.Set("Host", "fund.eastmoney.com")
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

// GetData 发送请求
func (c *EastMoneyFunds) JsObjectToJSON(body []byte) (map[string]interface{}, error) {
	content := string(body)
	content = strings.Replace(content, "var rankData = ", "", -1)
	content = strings.Replace(content, "};", "}", -1)

	keys := `datas:1,allRecords:1626,pageIndex:1,pageNum:1,allPages:1626,allNum:8309,gpNum:1626,hhNum:4518,zqNum:1984,zsNum:1175,bbNum:0,qdiiNum:181,etfNum:0,lofNum:330,fofNum:171`
	regs, _ := regexp.Compile(`:\d{1,}`)
	keys = regs.ReplaceAllString(keys, "")
	arr := strings.Split(keys, ",")
	for _, key := range arr {
		content = strings.Replace(content, key, fmt.Sprintf(`"%s"`, key), -1)
	}
	fundsjson := map[string]interface{}{}
	err := json.Unmarshal([]byte(content), &fundsjson)

	if err != nil {
		return nil, err
	}
	return fundsjson, nil
}
