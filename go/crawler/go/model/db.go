package model

import "gorm.io/gorm"
var DB*gorm.DB
//主表
type Info struct {
	gorm.Model
	Url      string `gorm:"unique;type:text"`
	Name     string
	Version  string 
	Details  []Detail `gorm:"foreignKey:IndexId"`
	LastUpdate string 
}

//副表
type Detail struct {
	gorm.Model
	IndexId   uint `gorm:"index"`
	Guid     string `gorm:"uninqueIndex"`
	Title     string
	Link      string `gorm:"type:text"`
	AutoLink  string
	PublishedTime int64
}
