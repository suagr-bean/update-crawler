package service

import (
	"project/model"
	"project/pkg"
	"testing"
)

func TestWork(t *testing.T) {
	path := "../data/rss.db"
	db, _ := pkg.DBInit(path)
	model.DB = db
	url := "https://wiwi.blog/blog/rss.xml"
	WorkService(url)
}
