package dao

import ("time"
   "project/model"
)
//根据时间查询待爬的
func QueryNextTime( times time.Time) []model.Info{
	var info []model.Info
	model.DB.Where("next_crawler_time<=?",times).Find(&info)
	return info
}