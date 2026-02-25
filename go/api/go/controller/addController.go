package controller

import (
	"encoding/json"
	"net/http"
	"project/dao"
	"project/model"
	"project/service"
)

// 添加URL 到数据库
func AddController(resp http.ResponseWriter, req *http.Request) {
	var data model.Data
	var result model.Result
	resp.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		result = model.Result{
			Code:    400,
			Message: "参数不对",
		}
		json.NewEncoder(resp).Encode(&result)
		return
	}
	//先查数据库有无数据
	deal, err := dao.QueryData(data.Url)
	if err != nil {
		result = model.Result{
			Code:    500,
			Message: "服务器内部错误",
		}
		json.NewEncoder(resp).Encode(&result)
		return
	}
	if deal == true {
		result = model.Result{
			Code:    409,
			Message: "数据库存在URL",
		}
	} else {
		s := service.AddService(data.Url)
		if s == true {
			result = model.Result{
				Code:    200,
				Message: "保存成功",
			}
		} else {
			result = model.Result{
				Code:    500,
				Message: "保存失败服务器内部错误",
			}
		}
	}

	json.NewEncoder(resp).Encode(&result)
}
