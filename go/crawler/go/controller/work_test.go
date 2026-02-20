package controller

import (
	"fmt"
	"project/dao"
	"testing"
)

func TestWork(t *testing.T) {
	path := "../data/test.db"
	dao.Init(path)
	WorkController()
	fmt.Println("end")
}
