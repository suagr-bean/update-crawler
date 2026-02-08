package dataDB

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBData struct {
	gorm.Model
	Name       string   `gorm:"type:string;unique;notnull"`
	Url        string   `gorm:"type:text;unique;notnull"`
	Version    string   `gorm:"type:string"`
	LastUpdate string   `gorm:"type:text"`
	Details    []Detail `gorm:"foreignKey:DBDataId"`
}

var db *gorm.DB

// 初始化数据库
func InitDB() error {
	path := "/workspaces/update-crawler/go/data/xml.db"

	var err error
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		fmt.Println("wrong")
		return err
	}
	err = db.AutoMigrate(&DBData{}, &Detail{})
	if err != nil {
		return err
	}
	return nil
}

type Detail struct {
	gorm.Model
	DBDataId uint `gorm:"index"`
	Title    string
	Link     string `gorm:"type:text"`
	Path     string
}
