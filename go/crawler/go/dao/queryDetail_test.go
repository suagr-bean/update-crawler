package dao

import (
	"fmt"
	"project/model"

	"testing"
)
func TestQuery(t*testing.T){
	path:="../data/test.db"
	err:=Init(path)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	show:=model.Show{}
	show.Url="https://feeds.simplecast.com/dHoohVNH"
	show.Size=2
	
	result,err:=QueryDetail(show)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	for _,v:=range result{
		fmt.Printf("标题%v\n链接%v\n",v.Title,v.Link)
	}
}