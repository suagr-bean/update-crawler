package service

import (
	"project/dao"
	"project/detect"
	"project/model"
	"fmt"
)
func AddService(url string)bool{
	fmt.Printf("addservice")
	result,err:=detect.Detect(url)
	if err!=nil{
      fmt.Println(err)
	  return  false
	} 
	var d []model.Detail
	for _,v:=range result.Articles{
     detail:= model.Detail{
       Title:v.Title,
	   Link:v.Link,
	   AutoLink:v.AutoLink,
	  }
	  d=append(d,detail)
	}
	info :=model.Info{
     Name:result.Name,
     Version:result.Version,
	 LastUpdate:result.LastUpdate,
	 Url:url,
     Details:d,
	}
	err=dao.Create(&info)
	if err!=nil{
      
      return false 
	}
	return true
	
}