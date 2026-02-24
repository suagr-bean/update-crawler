package service

import (
	"fmt"
	"project/Utils"
	"project/dao"
	"project/detect"
	"project/model"
	"time"
)

// 检查更新
func WorkService(url string) error {
	dealData, err := detect.Detect(url)
	if err != nil {
		return err
	}
	show := model.Show{
		Url:   url,
		Start: 0,
		Size:  1,
	}

	query, _ := dao.QueryDetail(show)
	interval := query.Info.(model.Info).DoMinute
	check := query.Details[0].Guid
	var detail []model.Detail
	for _, v := range dealData.Articles {
		if v.Guid == check {
			break
		}
		d := model.Detail{
			Title:         v.Title,
			Link:          v.Link,
			AutoLink:      v.AutoLink,
			Guid:          v.Guid,
			PublishedTime: v.PublishTime,
		}
		detail = append(detail, d)
	}
	//更新

	if len(detail) > 0 {
		newinter := int(float64(interval) * 0.9)
		next := Utils.DealTime(newinter)
		info := model.Info{
			Url:             url,
			Details:         detail,
			LastUpdate:      detail[0].Title,
			CrawlerTime:     time.Now(),
			NextCrawlerTime: next,
			DoMinute:        newinter,
		}
		dao.UpdateInfo(info)
		fmt.Println("已更新")
	} else {
		newinter := int(float64(interval) * 1.1)
		next := Utils.DealTime(newinter)
		info := model.Info{
			Url:             url,
			CrawlerTime:     time.Now(),
			NextCrawlerTime: next,
			DoMinute:        newinter,
		}
		dao.UpdateInfo(info)
		fmt.Println("没更新")
	}
	return nil
}
