package service

import (
	"fmt"
	"project/dao"
	"project/model"

)

//处理查询主表
func ShowService(show model.Show) ([]model.InfoResult,error){
	result, err := dao.QueryUrl(show)
	if err != nil {
		
		return []model.InfoResult{},fmt.Errorf("dao.QueryURl wrong:%w info:showstart:%wshowSize%w",err,show.Size,show.Start)
	}

	var r []model.InfoResult

	for _, v := range result {
		showresult := model.InfoResult{
			Name:       v.Name,
			LastUpdate: v.LastUpdate,
			Url:        v.Url,
			
		}
		r = append(r, showresult)
	}
	return r,nil
}
