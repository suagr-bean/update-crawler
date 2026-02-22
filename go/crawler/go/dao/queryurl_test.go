package dao

import (
	"project/model"
	"testing"
)
func TestQueryUrl(t*testing.T){
	path:="../data/test.db"
	Init(path)
	show:=model.Show{
		Size:50,
		Start:2,
	}
	query,_:=QueryUrl(show)
	t.Logf("数据%v",len(query))
}