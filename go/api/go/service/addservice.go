package service

import (
	"project/Utils"
	"project/dao"
	"project/detect"
	"project/model"
	"time"
	"fmt"
)

// 添加URL
func AddService(url string) bool {
	//爬取
	 con:=model.Context{
		Url:url,
	 }
	result, err := detect.Detect(con)
	if err != nil {
		return false
	}
	//健壮性 
    if result.Name==""{
		fmt.Println("出的数据是空")
    return false 
	}
     etag:=result.Etag
	 lastModified:=result.LastModified
	//组装文章
	var d []model.Detail
	for _, v := range result.Articles {
		detail := model.Detail{
			Title:         v.Title,
			Link:          v.Link,
			AutoLink:      v.AutoLink,
			PublishedTime: v.PublishTime,
			Guid:          v.Guid,
		}

		d = append(d, detail)
	}
	//content 1代表播客 0代表文章
	var content int
	var times int

	//判断有无音频
	if result.Articles[0].AutoLink != "" {
		content = 1
		times = 1440

	} else {
		content = 0
		times = 30
	}

	next := Utils.DealTime(times)
	info := model.Info{
		Name:            result.Name,
		Version:         result.Version,
		LastUpdate:      result.LastUpdate,
		Url:             url,
		Details:         d,
		ContentType:     content,
		DoMinute:        times,
		CrawlerTime:     time.Now(),
		NextCrawlerTime: next,
		LastModified:lastModified,
		Etag:etag,
	}

	//存进数据库
	err = dao.Create(&info)
	if err != nil {
		return false
	}
	return true
}
