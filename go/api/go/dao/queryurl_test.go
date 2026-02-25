package dao

import (
	"project/model"
	"project/pkg"
	"testing"
)

func TestQueryUrl(t *testing.T) {
	path := "../data/test.db"
	db, _ := pkg.DBInit(path)
	model.DB = db
	show := model.Show{
		Size:  50,
		Start: 2,
	}
	query, _ := QueryUrl(show)
	t.Logf("数据%v", len(query.Info.([]model.Info)))
	t.Logf("count%v", query.Count)
}
