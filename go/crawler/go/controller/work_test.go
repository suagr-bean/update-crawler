package controller

import (

	"project/dao"
	"testing"
)

func TestWork(t *testing.T) {
	path := "../data/test.db"
	dao.Init(path)
	WorkController()

}
