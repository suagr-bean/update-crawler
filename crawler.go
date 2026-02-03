package main

import (
	"database/sql"
	"fmt"
	"io"

	"net/http"
)
type Data struct {
	Url string 
	Etag string
	Code int
}
func start(db*sql.DB){
	rows,_:=db.Query("SELECT url FROM crawler")
	for rows.Next(){
		var url string
		rows.Scan(&url)
       fmt.Println(url)
	  data,err:= Crawler(url,"")
	  if err!=nil{
		fmt.Println("wrong")
	  }
	  fmt.Println(data)
	}
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
   datas,_:=io.ReadAll(resp.Body)
   fmt.Println(string(datas))
  fmt.Println(resp.StatusCode)
  
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