package main

import (
	"fmt"
	"strings"

	"net/http"
)

// 尝试请求头 看看需要用哪种策略
func HeaderTry(task []Task) []Task {
	ua := "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Mobile Safari/537.36"
	accept := "*/*"
	client := &http.Client{}
	for i := 0; i < len(task); i++ {
		url := task[i].Url
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("wrong")
			continue
		}

		req.Header.Set("User-Agent", ua)
		req.Header.Set("Accept", accept)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("错误")
			continue
		}
		_ = resp.StatusCode
		//请求头清洗 确实策略
		etag := resp.Header.Get("etag")

		if etag != "" { //有etag
			if strings.HasPrefix(etag, "W/") { //弱etag 放弃选择
				task[i].Etag = ""
				fmt.Println("弱etag", task[i].Etag)
			} else {
				task[i].Etag = etag
			}
		}
		if etag == "" {
			last := resp.Header.Get("Last-Modified")
			task[i].Last = last
		}
		//如果没有 特殊处理 存body 哈希

		resp.Body.Close()
	}
	return task
}
