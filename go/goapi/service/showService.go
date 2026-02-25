package service

import (
	"fmt"
	"project/dao"
	"project/model"
)

// 处理查询主表
func ShowService(show model.Show) ([]model.InfoResult, error) {
	result, err := dao.QueryUrl(show)
	if err != nil {

		return []model.InfoResult{}, fmt.Errorf("dao.QueryURl wrong:%w info:showstart:%dshowSize%d", err, show.Size, show.Start)
	}

	var r []model.InfoResult

	for _, v := range result.Info.([]model.Info) {
		var contextType string
		if v.ContentType == 0 {
			contextType = "article"
		} else if v.ContentType == 1 {
			contextType = "podcast"
		}
		showresult := model.InfoResult{
			Name:        v.Name,
			LastUpdate:  v.LastUpdate,
			Url:         v.Url,
			Count:       result.Count,
			ContextType: contextType,
		}
		r = append(r, showresult)
	}
	return r, nil
}
