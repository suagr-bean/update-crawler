package main

import (
	"fmt"
	"project/controller"

	"project/model"
	"project/pkg"
)

func main() {
	fmt.Println("启动")
	//打开数据库
	path := "data/test2.db"
	db, err := pkg.DBInit(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	model.DB = db
	addr := "127.0.0.1:6379"
	err = pkg.RedisInit(addr, "", 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	go controller.Do()
	//爬虫工作 设置工作时间
	// 	go controller.WorkStart(10*time.Hour)
	//路由

	controller.Route()
}
