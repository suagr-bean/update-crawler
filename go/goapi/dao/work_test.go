package dao

import (
	"project/pkg"
	"testing"
)
func TestWork(t*testing.T){
	path:="/home/user/rssread/go/goapi/data/test3.db"
	pkg.DBInit(path)
	result:=QueryWork(1)
	t.Log(result.Details[0].Guid)
}