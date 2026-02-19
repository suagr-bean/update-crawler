package detect

import (
	"net/http"
	"project/model"
)

// 抓取
func Crawler(url string) (model.Context, error) {
	client := http.Client{}
   context:= model.Context{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return context, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/xml,text/xml,*/*")
	resp, err := client.Do(req)
	if err != nil {
		return context, err
	}
	
	return model.Context{
		Url:  url,
		Code: resp.StatusCode,
		Body: resp.Body,
	}, nil
}
