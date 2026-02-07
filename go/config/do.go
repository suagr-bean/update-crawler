package config

import (
	
	"project/dataDB"
	"project/detect"
	"fmt"
	"log"
)
type  dealing struct{
     rss string
	 name string 
	 url string
     
}
//读配置文件和数据库内容比对
func Do(path string) {
     //读配置文件
	result,err:=ReadFile(path)
	if err!=nil{
       log.Printf("read file err%v",err)
	}
	 //dataDB.InitDB()//初始化数据库
     
	//遍历URL
	for _,data:=range result{
       name:=data.Name
	   //检查数据是否存在数据库
	   deal,err:=dataDB.Query(name)
	   if err!=nil{
		log.Printf("query err%v",err)
	   }
	   //存在
	   if deal.Name==name{
		fmt.Println("数据库存在数据")
		 continue
	   }else {
		//不存在 探测版本并记录数据库
		 result,err:=detect.Detect(data.Url)
		 if err!=nil{
			log.Printf("detect%v",err)
		 }
		 newdata:=dataDB.DBData{
			Name: data.Name,
			Url:   data.Url,
			Version: result.Version,
			LastUpdate :result.Lastupdate,
			}
			err=newdata.Create()//存进数据库
			if err!=nil{
               log.Printf("create err%v",err)
			}
	   }
	}
}