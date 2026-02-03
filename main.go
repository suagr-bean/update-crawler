package main

import "time"
func main(){
	db,err:=ReadSQL()//打开数据库
	if err!=nil{
		db.Close()
	}
	for {
	SelectData(db)
	start(db)
	time.Sleep(15*time.Minute)
	}
	/*
	for{
	 fmt.Println("开始")
     
	
	 for i,v:=range urls{
	   cleanUrl:= strings.TrimSpace(v)
	   fmt.Println(cleanUrl)
	   err:=EtagAndSave(cleanUrl,i)
	   if err!=true{
        fmt.Println(i,"失败")
		continue
	   }
	   DealCmd(cleanUrl)
	 }
	 fmt.Println("结束")
	 time.Sleep(15*time.Second)
	 
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
	*/
}