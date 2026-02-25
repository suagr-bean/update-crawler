package pkg

import (
	"errors"

	"project/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// 数据库初始化
func DBInit(path string) (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Info{}, &model.Detail{})
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("db is nil")
	}
	return db, nil

}
