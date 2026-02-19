package controller

import (
	"project/dao"
	"testing"
	"fmt"
)

func TestWork(t *testing.T){
	path:="../data/test.db"
	dao.Init(path)
	WorkController()
	fmt.Println("end")
}