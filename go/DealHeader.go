package main

import "net/http"

type CrawlerHeader struct {
	Url      string
	Etag     string
	Modified string
	Strate   string
	Update   bool
}

// 请求头策略
func DealHeader(c CrawlerHeader) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, nil)
	if err != nil {

	}
	switch c.Strate {

	case "Etag":
		req.Header.Set("If-None-Match", c.Etag)
	case "Modified":
		req.Header.Set("If-Modified-Since", c.Modified)
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotModified {
		c.Update = false
	} else {
		if c.Strate == "Etag" {
			newEtag := resp.Header.Get("Etag")
			c.Etag = newEtag
			c.Update = true
		} else {
			modified := resp.Header.Get("Last-Modified")
			c.Modified = modified
			c.Update = true
		}
	}
}
