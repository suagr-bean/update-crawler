package main

import (
	"fmt"
	"os"
	"strings"
)
//读取需要爬的网站
func ReadFile()[]string {
	url,err:=os.ReadFile("url.text")
	if err!=nil{
		fmt.Println("错误，读不到文件")
	}
	urls:=strings.Split(string(url),"\n")
	return urls
}
