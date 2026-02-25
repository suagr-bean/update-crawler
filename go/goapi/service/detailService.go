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
	if len(result.Details) == 0 {
		return detail, nil
	}
	for _, v := range result.Details {
		d := model.DetailData{
			Title:         v.Title,
			Link:          v.Link,
			AutoLink:      v.AutoLink,
			PublishedTime: v.PublishedTime,
			Count:         result.Count,
		}
		detail = append(detail, d)
	}
	return detail, nil
}
