package controller

import (
	"net/http"
	"project/dao"
	"project/model"
	"strconv"
)

// 模糊查询
func SeachController(resp http.ResponseWriter, req *http.Request) {
	var show model.Show
	getUrl := req.URL.Query().Get("u")
	getTitle := req.URL.Query().Get("t")
	Startstr := req.URL.Query().Get("st")
	Sizestr := req.URL.Query().Get("si")
	start, _ := strconv.Atoi(Startstr)
	size, _ := strconv.Atoi(Sizestr)
	if getUrl == "" {
		return
	}
	if getTitle == "" {
		return
	}
	show.Url = getUrl
	show.SeachTitle = getTitle
	if show.Size >= 0 {
		show.Size = size
	}
	if show.Start >= 0 {
		show.Start = start
	}

	dao.QuerySeach(show)
}
