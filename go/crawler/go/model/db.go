package model

import "gorm.io/gorm"
import "time"

var DB *gorm.DB

// 主表
type Info struct {
	gorm.Model
	Url             string `gorm:"unique;type:text"`
	Name            string
	Version         string
	Details         []Detail `gorm:"foreignKey:IndexId"`
	LastUpdate      string
	ContentType     int //0 文章 1播客
	DoMinute        int
	CrawlerTime     time.Time
	NextCrawlerTime time.Time `gorm:"index"`
	LastModified string 
	Etag string 
}

// 副表
type Detail struct {
	gorm.Model
	IndexId       uint   `gorm:"index"`
	Guid          string `gorm:"uninqueIndex"`
	Title         string
	Link          string `gorm:"type:text"`
	AutoLink      string
	PublishedTime int64 `gorm:"index"`
}
