package main

import (
	"database/sql"
	"strings"
	"fmt"
)
//初始化读 URL 保存数据库
func SelectData(db*sql.DB){
 Urls:=ReadFile()
 query:="SELECT COUNT(*) FROM crawler WHERE url = ?"
  count:=0
  for _,v:=range Urls{
	cleanurl:=strings.TrimSpace(v)
   err:=db.QueryRow(query,cleanurl).Scan(&count)
    if err!=nil{
       
	}
	if count==0{
		strate:=Detect(cleanurl)
		fmt.Println(strate)
		if strate==404{
			continue
		}
	  _,err:=db.Exec("INSERT INTO crawler(url,strategy) VALUES(?,?)",cleanurl,strate)
       if err!=nil{
		fmt.Println("插入数据失败")
	   }
	}
  }
  
}