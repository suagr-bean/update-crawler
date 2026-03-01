package service

import (
	
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

	//判断有无音频
	if result.Articles[0].AutoLink != "" {
		content = 1
	} else {
		content = 0	
	}
	info := model.Info{
		Name:            result.Name,
		Version:         result.Version,
		LastUpdate:      result.LastUpdate,
		Url:             url,
		Details:         d,
		ContentType:     content,
		CrawlerTime:     time.Now().Unix(),
		LastModified:lastModified,
		Etag:etag,
	}

	//存进数据库
	err = dao.Create(&info)
	if err != nil {
		return false
	}
	//处理高频时间 存进数据库
   frequency:= Caltimes(url)
    next:=CalDo(frequency)
	dao.UpdateDoNext(url,frequency,next)
	return true
}
