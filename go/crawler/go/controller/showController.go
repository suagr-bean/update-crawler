package controller

import (
	"encoding/json"
	"net/http"
	"project/model"
	"project/service"
	"strconv"
)
//查询主表
func ShowController(resp http.ResponseWriter, req *http.Request) {
     var result model.Result
	query := req.URL.Query()
	startstr := query.Get("start")
	sizestr := query.Get("size")
	start, _ := strconv.Atoi(startstr)
	size, _ := strconv.Atoi(sizestr)
	if size == 0 {
		size = 15 
	}
	if start==0{
		start=1
	}
	show := model.Show{
		Size:  size,
		Start: start,
	}
	infoResult,err:= service.ShowService(show)
	if err!=nil{
	  return 
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	result.Code=200
	result.Info=infoResult
	result.Message="successful"
	json.NewEncoder(resp).Encode(&result)
}
