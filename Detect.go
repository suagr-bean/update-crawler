package main

 import( "net/http"
   "fmt")

//探测 网站有没有防护
func Detect(url string)int{
   client:= &http.Client{}
   req,err:=client.Head(url)
   if err!=nil{
	fmt.Println("爬请求头错误")
   }
   state:=0
   switch req.StatusCode{
   case 200://简单
   state=200
   case 404: //网站写错
   state =404 
   case 403: //复杂
   state =403
   case 503: // 服务器问题
	state =503
   }
   return state
}