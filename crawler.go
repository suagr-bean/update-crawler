package main

import (
	"net/http"
	"os"
	"fmt"
)
type Data struct {
	Url string 
	Etag string
}
func EtagAndSave(url string)bool{
	 readetag,_:=os.ReadFile("etag.text")
	data,err:= Crawler(url,string(readetag))
	if err!=nil{
		fmt.Println("爬取失败")
		return false 
	}
	if data.Etag!=""{
     os.WriteFile("etag.text",[]byte(data.Etag),0644)
	}else{
		fmt.Println("没更新")
		return false
	}
	return true
}
func Crawler(url string,etag string)(Data,error){
   client:=&http.Client{}
   data:=Data{}
   req,_:= http.NewRequest("GET",url,nil)
   req.Header.Set("If-None-Match",etag)
   resp,err:=client.Do(req)
  if err!=nil{
	return data,err
  }
  defer resp.Body.Close()
  update:=false
  switch resp.StatusCode{
  
  case http.StatusNotModified:
	 update=false
  case http.StatusOK:
	update=true 
  }
  if update==true{
	etag:=resp.Header.Get("Etag")
	if etag!=""{
	 data:=Data{
      Url:url,
      Etag: etag,
	}
   
	return data,err
} 
  }
  return data,err
}