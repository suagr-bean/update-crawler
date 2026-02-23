package service

import (
	"project/dao"
	"project/detect"
	"project/model"
	"time"
)

//添加URL
func AddService(url string) bool {
   //爬取
	result, err := detect.Detect(url)
	if err != nil {
		return false
	}
    
	//组装文章
	var d []model.Detail
	for _, v := range result.Articles {
		detail := model.Detail{
			Title:         v.Title,
			Link:          v.Link,
			AutoLink:      v.AutoLink,
			PublishedTime: v.PublishTime,
			Guid :v.Guid,
		}
        
		d = append(d, detail)
	}
	 //content 1代表播客 0代表文章
	 var content int 
	 var times int
	 var duration time.Duration
	 //判断有无音频 
     if result.Articles[0].AutoLink!=""{
        content=1
		times=1440
       
	 }else{
		content=0
		times=60
	 }
     duration=time.Duration(times)*time.Minute
	 next:=time.Now().Add(duration)
	info := model.Info{
		Name:       result.Name,
		Version:    result.Version,
		LastUpdate: result.LastUpdate,
		Url:        url,
		Details:    d,
		ContentType :content,
		DoMinute:  times,
		CrawlerTime: time.Now(),
		NextCrawlerTime: next,
	}

	//存进数据库
	err = dao.Create(&info)
	if err != nil {
		return false
	}
	return true
}
