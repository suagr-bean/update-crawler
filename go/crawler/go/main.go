package main

import (
	"fmt"
	"project/controller"
	
	"project/model"
	"project/pkg"
	"time"
)

func main() {
	fmt.Println("启动")
	//打开数据库
	path := "data/test2.db"
	db,err:=pkg.DBInit(path)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	model.DB=db
//爬虫工作 设置工作时间
   	go controller.WorkStart(10*time.Hour)
	//路由
	 
	controller.Route()
}
