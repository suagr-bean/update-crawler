package dao

import (
	
	"project/model"
	"testing"
)

func TestDelete(t *testing.T){
	Path:="../data/rss.db"
    Init(Path)
	model.DB.Where("id=?",1376).Delete(&model.Detail{})
}