package crawler

import (
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
)

// XueqiuCookie 取得雪球网的cookie  同时取几个首页数据
type XueqiuCookie struct {
	HomePageData []map[string]string
}

func (t *XueqiuCookie) Get() []*http.Cookie {
	var jar *cookiejar.Jar
	jar, _ = cookiejar.New(nil)
	httpClient := &http.Client{
		Jar: jar,
	}
	site := "https://xueqiu.com"
	var httpReq *http.Request
	httpReq, _ = http.NewRequest("GET", site, nil)
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		log.Println(err.Error())
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		log.Printf("status code error: %d %s", httpResp.StatusCode, httpResp.Status)
	}

	// Load the HTML document
	// reader := simplifiedchinese.GB18030.NewDecoder().Reader(httpResp.Body)
	doc, err := goquery.NewDocumentFromReader(httpResp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	maps := []map[string]string{}
	doc.Find("h3").Each(func(i int, cont *goquery.Selection) {
		one := map[string]string{}
		one["title"] = cont.Text()
		one["content"] = cont.Next().Text()
		one["link"] = site + cont.Prev().AttrOr("href", "")
		maps = append(maps, one)
	})

	t.HomePageData = maps

	cookies := jar.Cookies(httpReq.URL)
	return cookies
}
