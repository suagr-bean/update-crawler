package service

import (
	"project/dao"
	"project/detect"
	"project/model"
)

func AddService(url string) bool {
	result, err := detect.Detect(url)
	if err != nil {
		return false
	}
	//组装文章
	var d []model.Detail
	for _, v := range result.Articles {
		detail := model.Detail{
			Title:         v.Title,
			Link:          v.Link,
			AutoLink:      v.AutoLink,
			PublishedTime: v.PublishTime,
			Guid :v.Guid,
		}
		d = append(d, detail)
	}
	info := model.Info{
		Name:       result.Name,
		Version:    result.Version,
		LastUpdate: result.LastUpdate,
		Url:        url,
		Details:    d,
	}

	//存进数据库
	err = dao.Create(&info)
	if err != nil {
		return false
	}
	return true
}
