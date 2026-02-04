package main

// 策略爬虫
func CrawlerDeal(data Data) string {
	c := CrawlerHeader{}
	c.Url = data.Url
	if data.Etag != "" {
		c.Etag = data.Etag
		c.Strate = "ETag"
	} else if data.Modified != "" {
		c.Modified = data.Modified
		c.Strate = "Modified"
	}
	DealHeader(c)
	if c.Update == true {
		if c.Strate == "Etag" {
			return c.Etag
		} else {
			return c.Modified
		}
	}
	return ""
}
