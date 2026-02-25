package dao

import (
	"project/model"
	"time"
)

// 根据时间查询待爬的
func QueryNextTime(times time.Time) []model.Info {
	var info []model.Info
	model.DB.Where("next_crawler_time<=?", times).Find(&info)
	return info
}

// 查询
func QueryWork(id int) model.Info {
	info := model.Info{}
	model.DB.Where("id=?", id).First(&info)
	return info
}
