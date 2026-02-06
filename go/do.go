package main

import (
	"project/config"
	"project/dataDB"
	"project/detect"
	"fmt"
)
type  dealing struct{
     rss string
	 name string 
	 url string
     
}
//处理流程
func Do(path string) {
	result,err:=config.ReadFile(path)
	if err!=nil{

	}
	 dataDB.InitDB()//初始化数据库
     
	//遍历URL
	for _,data:=range result{
       name:=data.Name
	   //检查数据是否存在数据库
	   deal,err:=dataDB.Query(name)
	   if err!=nil{}
	   //存在
	   if deal{
		continue
	   }else {
		 version,err:=detect.Detect(data.Url)
		 if err!=nil{}
		 fmt.Println("版本",version)
		 newdata:=dataDB.DBData{
			Name: data.Name,
			Url:   data.Url,
			Version: version,
			}
			err=newdata.Create()
			if err!=nil{}
	   }
	}
}