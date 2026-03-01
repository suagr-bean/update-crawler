package service

import (
	"fmt"
	
	"project/dao"
	"project/detect"
	"project/model"
	"time"
)

// 检查更新
func WorkService( index int) error {
	fmt.Println(index)
	q:=dao.QueryWork(index)
	con:=model.Context{
		Etag:q.Etag,
		LastModified:q.LastModified,
		Url:q.Url,
       Version:q.Version,
	}
	 shows:=model.Show{
       Start:1,
	   Size:1,
	   Url:q.Url,
	 }
    d,_:=dao.QueryDetail(shows)
	dealData, err := detect.Detect(con)
	if err != nil {
		return err
	}
	 cal:= model.TimeCal{
		Interval:q.DoMinute,
		IsUpdate: false ,
       Try :q.Try,
	 }
	var info model.Info
	//处理304情况
	
	if dealData.Etag==""{
		do:=DoTime(cal)
	  info.Url=q.Url
      info.DoMinute=do.Interval
	  info.CrawlerTime=time.Now().Unix()
	  info.NextCrawlerTime=time.Now().Add(1*time.Hour).Unix()
	  info.Try=q.Try+1

       fmt.Println("304 没更新")
	}else {
		//和数据库存的guid比较
	check :=d.Details[0].Guid
	var detail []model.Detail
	for _, v := range dealData.Articles {
		if v.Guid == check {
			break
		}
		d := model.Detail{
			Title:         v.Title,
			Link:          v.Link,
			AutoLink:      v.AutoLink,
			Guid:          v.Guid,
			PublishedTime: v.PublishTime,
		}
		detail = append(detail, d)
	}
	//更新

	if len(detail) > 0 {
		cal.IsUpdate=true
		do:=DoTime(cal)
		info = model.Info{
			Url:             con.Url,
			Details:         detail,
			LastUpdate:      detail[0].Title,
			CrawlerTime:     time.Now().Unix(),
			NextCrawlerTime: time.Now().Add(24*time.Hour).Unix(),
			DoMinute:        do.Interval,
			LastModified: dealData.LastModified,
			Etag: dealData.Etag,
			Try:0,
		}
		fmt.Println("已更新")
	} else {
		do:=DoTime(cal)
		var next int64
		 var try int
		//重试次数
         if do.Try<3{
           next=time.Now().Add(1*time.Hour).Unix()
		   try=do.Try+1
		 }else{
		   next=time.Now().Add(20*time.Hour).Unix()
            try=0
		 }
		 
		info = model.Info{
			Url:             con.Url,
			CrawlerTime:     time.Now().Unix(),
			NextCrawlerTime: next,
			DoMinute:        do.Interval,
			LastModified: dealData.LastModified,
			Etag:dealData.Etag,
			Try: try,
		}
		fmt.Println("没更新")
	}
}
	 dao.UpdateInfo(info)
	return nil
}
