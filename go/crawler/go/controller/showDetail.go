package controller

import (
	"encoding/json"
	"net/http"
	"project/model"
	"project/service"
	
)
func ShowDetail(resp http.ResponseWriter, req*http.Request){
	  
	 var r model.Result
      resp.Header().Set("Content-Type","application/json")
	 url:= req.URL.Query().Get("url")
	 
	 result,err:=service.DetailService(url)
	 if err!=nil{
		return 
	 }
     r=model.Result{
		Code:200,
		Message:"查询成功",
		Data:result,
	 }
	 json.NewEncoder(resp).Encode(&r)
}