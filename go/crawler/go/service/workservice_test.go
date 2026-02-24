package service

import (
	"project/model"
	"project/pkg"
	"testing"
)

func TestWork(t *testing.T) {
	path := "/home/user/rssread/go/crawler/go/data/test2.db"
	db, _ := pkg.DBInit(path)
	model.DB = db
	url := "https://wiwi.blog/blog/rss.xml"
	 err:=WorkService(url)
	 if err!=nil{
		t.Log(err)
	 }
}
