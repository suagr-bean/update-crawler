package main

import (
	"project/dataDB"
	"project/detect"
	"fmt"
)

func listen(){
	data,err:=dataDB.QueryAll()
    if err!=nil{
      fmt.Println("err")
	}
      Loop(data) 
}
func Loop(data []dataDB.DBData){
    for _,v:=range data{
	  version:=v.Version
	  context,err:= detect.CrawlerReq(v.Url)
	 if err!=nil{
		fmt.Printf("流错误",err)
		return 
	 }
	  switch version{
	  case "RSS2.0":
	   rss:=&detect.Rss{}
        rss.AccessInfo(context)
        check:=rss.UpdateTitle()
		if check==v.LastUpdate{
			fmt.Println("没更新")
		}else{
			fmt.Println("已更新")
		}
	  }
	}
}