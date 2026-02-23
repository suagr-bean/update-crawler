package dao

import (
	"project/model"
	"project/pkg"
	"testing"
)
func TestQuerySeach(t*testing.T){
	path:="../data/test.db"
	db,err:=pkg.DBInit(path)
	model.DB=db
	if err!=nil{
		t.Logf("错误%v",err)
	}
	var show model.Show
	show.Url="https://feeds.megaphone.fm/hubermanlab"
	show.SeachTitle="sleep"
	show.Start=1
	show.Size=2
	result,err:=QuerySeach(show)
	if err!=nil{
		t.Logf("错误原因%v",err)
		return 
	}
	for _,v:=range result{
		t.Logf("标题%v\n",v.Title)
	}
}