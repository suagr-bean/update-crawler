package service

import (
	"project/model"
	"project/pkg"
	"testing"
)

func TestCalTimes(t *testing.T){
	path:="/home/user/rssread/go/goapi/data/test3.db"
	db,err:=pkg.DBInit(path)
	model.DB=db
	if err!=nil{
		t.Logf("数据库初始化错误%v",err)
		return 
	}
	url:="https://wiwi.blog/blog/rss.xml"
	result:=Caltimes(url)
	deal:=CalDo(result)
	t.Log("下次爬取时间",deal)
}