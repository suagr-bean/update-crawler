package dao

import (
	"project/model"
	"project/pkg"
	"strconv"
	"testing"
)

func TestCreate(t *testing.T) {
	path := "../data/test.db"
	db, _ := pkg.DBInit(path)
	model.DB = db
	for i := 0; i < 10000; i++ {
		var info model.Info
		str := strconv.Itoa(i)
		info.Url = "http://test-performance.com/item/" + str
		Create(&info)
	}
}
