package controller

import (
	
	"project/model"
	"project/pkg"
	"testing"
)

func TestWork(t *testing.T) {
	path := "../data/test.db"
	db,_:=pkg.DBInit(path)
 model.DB=db
	WorkController()

}
