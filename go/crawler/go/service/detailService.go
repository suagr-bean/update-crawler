package service

import (
	"project/dao"
	"project/model"
)

func DetailService(show model.Show) ([]model.DetailData, error) {
	var detail []model.DetailData
	result, err := dao.QueryDetail(show)
	if err != nil {
		return detail, err
	}
	for _, v := range result {
		d := model.DetailData{
			Title:    v.Title,
			Link:     v.Link,
			AutoLink: v.AutoLink,
            PublishedTime: v.PublishedTime,
		}
		detail = append(detail, d)
	}
	return detail, nil
}
