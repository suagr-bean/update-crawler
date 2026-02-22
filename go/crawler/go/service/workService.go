package service

import (
	"project/dao"
	"project/detect"
	"project/model"
    "fmt"
)
//检查更新 
func WorkService(url string){
	dealData,err:=detect.Detect(url)
    if err!=nil{
		return 
	}
	show:=model.Show{
		Url:url,
		Start:0,
		Size:1,
	}
	
	query,_:=dao.QueryDetail(show)
    check:=query[0].Guid
     var  detail []model.Detail
    for _,v:=range dealData.Articles{
        if v.Guid==check{
			break
		}
	   d:=model.Detail{
		 Title:v.Title,
		 Link:v.Link,
		 AutoLink:v.AutoLink,
		 Guid:v.Guid,
		 PublishedTime: v.PublishTime,
	   }
	   detail=append(detail,d)
	}
	 if len(detail)>0{
	 info:=model.Info{
		Name:dealData.Name,
		Version:dealData.Version,
		Url:url,
		Details:detail,
		LastUpdate:detail[0].Title,
	 }
	err:= dao.UpdateInfo(info)
	if err!=nil{
		return 
	}
      fmt.Println("更新")
	 }else{
	 fmt.Println("没更新")
	 }
}