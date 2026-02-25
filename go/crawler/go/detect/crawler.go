package detect

import (
	"io"
	"net/http"
	"project/model"
)

// 抓取
func Crawler( cont model.Context) (model.Context, error) {
	client := http.Client{}
   context:= model.Context{}
	req, err := http.NewRequest("GET", cont.Url, nil)
	if err != nil {
		return context, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/xml,text/xml,*/*")
	if cont.Etag!=""{
		req.Header.Set("If-None-Match",cont.Etag)
	}
	if cont.LastModified!=""{
		req.Header.Set("If-Modified-Since",cont.LastModified)
	}
	resp, err := client.Do(req)
	if err != nil {
		return context, err
	}
	defer resp.Body.Close()
	cont.Code=resp.StatusCode
	if cont.Code==http.StatusOK{
	    cont.Etag=resp.Header.Get("Etag")
		cont.LastModified=resp.Header.Get("Last-Modified")
		data,err:=io.ReadAll(resp.Body)
		if err!=nil{
			return cont,err 
		}
       cont.Body=data
	} 
	return cont ,nil
}
