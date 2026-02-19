package main

import (
	"fmt"
	"project/controller"
	"project/dao"
)

func main() {
	fmt.Println("启动")
	//打开数据库
	path := "data/rss.db"
	dao.Init(path)
	//路由
	controller.Route()

}
