package dao

import (
	"project/model"
	"project/pkg"
	"testing"
)

func TestDelete(t *testing.T) {
	Path := "/home/user/rssread/go/crawler/go/data/test2.db"
	db, _ := pkg.DBInit(Path)
	model.DB = db
	model.DB.Where("id=?", 1385).Delete(&model.Detail{})
}
