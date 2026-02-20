package main

import (
	"fmt"
	"project/controller"
	"project/dao"
	"time"
)

func main() {
	fmt.Println("启动")
	//打开数据库
	path := "data/rss.db"
	dao.Init(path)
//爬虫工作
	go controller.WorkStart(24*time.Hour)
	//路由
	controller.Route()
}
