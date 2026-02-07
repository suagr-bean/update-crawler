package main

import (
	"project/config"
	"project/dataDB"
)



func main(){
    dataDB.InitDB()//数据库连接
    path:="/workspaces/update-crawler/go/input.csv"
    config.Do(path)//读配置文件
    listen() 
}