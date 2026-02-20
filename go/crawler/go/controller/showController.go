package controller

import (
	"encoding/json"
	"net/http"
	"project/model"
	"project/service"
	"strconv"
)

func ShowController(resp http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	startstr := query.Get("start")
	sizestr := query.Get("size")
	start, _ := strconv.Atoi(startstr)
	size, _ := strconv.Atoi(sizestr)
	if size == 0 {
		size = 15 // Default to 10 if not specified
	}
	show := model.Show{
		Size:  size,
		Start: start,
	}
	deal := service.ShowService(show)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(&deal)

}
