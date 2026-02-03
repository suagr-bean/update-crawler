package main

import (
	"fmt"
	"os/exec"
	"time"
)
func main(){

	url:="https://wiwi.blog/blog/"
	for {
		fmt.Println("开始执行")
    deal:= EtagAndSave(url)
     if deal==true{
		DealCmd(url)
	 }
	 
	 time.Sleep(3*time.Minute)
	}
}
func DealCmd (url string){
	need:=fmt.Sprintf("termux-open %s",url)
  cmd:= exec.Command("termux-notification","--title","网站有更新","--action",need,
)
   err:=cmd.Run()
   if err!=nil{
    fmt.Println("通知失败")
   }
}