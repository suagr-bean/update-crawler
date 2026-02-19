package dao

import (
	"project/model"
	"strconv"
	"testing"
)

func TestCreate(t *testing.T) {
	path := "../data/test.db"
	Init(path)

	for i := 0; i < 10000; i++ {
		var info model.Info
		str := strconv.Itoa(i)
		info.Url = "http://test-performance.com/item/" + str
		Create(&info)
	}
}
