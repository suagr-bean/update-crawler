package controller

import (
	"encoding/json"
	"net/http"
	"project/model"
	"project/service"
	"strconv"
)

func ShowDetail(resp http.ResponseWriter, req *http.Request) {

	var r model.Result
	resp.Header().Set("Content-Type", "application/json")
	var show model.Show
	show.Url = req.URL.Query().Get("url")
	size := req.URL.Query().Get("size")
	start := req.URL.Query().Get("start")
	//转换Int
	SizeInt, err := strconv.Atoi(size)
	if err != nil {
		SizeInt = 0
	}
	StartInt, err := strconv.Atoi(start)
	if err != nil {
		StartInt = 0
	}
	show.Size = SizeInt
	show.Start = StartInt
	result, err := service.DetailService(show)
	if err != nil {
		return
	}
	r = model.Result{
		Code:    200,
		Message: "查询成功",
		Data:    result,
	}
	json.NewEncoder(resp).Encode(&r)
}
