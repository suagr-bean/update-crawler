package dao

import (
	"fmt"

	"testing"
)
func TestQuery(t*testing.T){
	path:="../data/test.db"
	err:=Init(path)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	url:="https://wiwi.blog/blog/rss.xml"
	result,err:=QueryDetail(url)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	for _,v:=range result{
		fmt.Printf("标题%v\n链接%v\n",v.Title,v.Link)
	}
}