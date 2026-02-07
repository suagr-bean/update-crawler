package detect

import (
	"io"
	"net/http"
)


//上下文连接
type Context struct {
	Code int
	Body io.ReadCloser
	URL  string
}

// 爬取
func CrawlerReq(url string) (Context, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Context{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/xml,text/xml,*/*")
	resp, err := client.Do(req)
	if err != nil {
		return Context{}, err
	}
	
	return Context{
		URL:  url,
		Code: resp.StatusCode,
		Body: resp.Body,
	}, nil
}
