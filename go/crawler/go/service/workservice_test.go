package service

import (
	"project/dao"
	"testing"
)

func TestWork(t *testing.T){
	path:="../data/rss.db"
    dao.Init(path)
	url:="https://wiwi.blog/blog/rss.xml"
	WorkService(url)
}