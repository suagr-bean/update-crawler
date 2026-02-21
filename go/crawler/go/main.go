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
	path := "data/test.db"
	dao.Init(path)
//爬虫工作 设置工作时间
	go controller.WorkStart(7*time.Hour)
	//路由
	controller.Route()
}
