package service

import (
	"project/dao"
	"project/model"
)
func DetailService(url string)([]model.DetailData,error){
	var detail [] model.DetailData
   result,err:= dao.QueryDetail(url)
   if err!=nil{
	return  detail,err
   }
   for _,v:=range result{
	 d:=model.DetailData{
		Title:v.Title,
		Link:v.Link,
		AutoLink:v.AutoLink,
	 }
	 detail=append(detail,d)
   }
    return detail,nil
}