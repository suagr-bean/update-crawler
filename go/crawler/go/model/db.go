package model

import (
 "gorm.io/gorm"
)
//目录表
type Info  struct {
	gorm.Model
	Name       string   `gorm:"type:string;unique;notnull"`
	Url        string   `gorm:"type:text;unique;notnull"`
	Version    string   `gorm:"type:string"`
	LastUpdate string   `gorm:"type:text"`
	Details    []Detail `gorm:"foreignKey:IndexId"`
}

var DB *gorm.DB

//副表
type Detail struct {
	gorm.Model
	IndexId uint `gorm:"index"`
	Title    string
	Link     string `gorm:"type:text"`
	AutoLink     string
}
