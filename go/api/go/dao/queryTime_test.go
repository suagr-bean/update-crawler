package dao

import (
	"project/model"
	"project/pkg"
	"testing"
	"time"
)

func TestQueryTime(t *testing.T) {
	path := "/home/user/rssread/go/crawler/go/data/test2.db"
	db, _ := pkg.DBInit(path)
	model.DB = db
	times := time.Now().Add(2 * time.Hour)
	t.Log(times)
	result := QueryNextTime(times)
	t.Logf("待爬数据%v", len(result))
}
