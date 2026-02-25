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
	index:=1
	 err:=WorkService(index)
	 if err!=nil{
		t.Log(err)
	 }
}
