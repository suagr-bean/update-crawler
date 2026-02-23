package dao

import (
	"project/model"
	"project/pkg"
	"testing"
)

func TestDelete(t *testing.T){
	Path:="../data/test2.db"
    db,_:=pkg.DBInit(Path)
	model.DB=db
	model.DB.Where("id=?",1,2).Delete(&model.Info{})
}