package crawler

import "net/http"

// XueqiuHeader 雪球伪装头部
type XueqiuHeader struct {
}

func (t XueqiuHeader) HqHeader() http.Header {

	header := http.Header{}
	// header.Add(":path", "")
	// header.Add("cookie", "")
	header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")

	return header
}
