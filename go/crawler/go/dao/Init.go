package dao

import (
	"project/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
 
//数据库初始化
func Init (path string)error{
	var err error 
	model.DB,err=gorm.Open(sqlite.Open(path),&gorm.Config{})
	if err!=nil{
		return err
	}
	err=model.DB.AutoMigrate(&model.Info{},&model.Detail{})
	if err!=nil{
		return err
	}
	return nil

}
