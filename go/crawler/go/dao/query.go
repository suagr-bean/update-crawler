package dao

import "project/model"
//查询主数据库
func QueryUrl(show model.Show)([]model.Info,error){
	var info [] model.Info
	if show.Size <= 0 {
		show.Size = 15
	}
	if show.Start < 0 {
		show.Start = 0
	}
	offset := show.Start * show.Size
	err := model.DB.Limit(show.Size).Offset(offset).Find(&info).Error
    if err != nil {
		return info, err
	}
	return info, nil
}
//查询数据库有无数据
func QueryData(url string)(bool,error){
	var info model.Info
	result:=model.DB.Limit(1).Where("url=?",url).Find(&info)
	if result.Error!=nil{
		return false,result.Error
	}
    if result.RowsAffected>0{
		return true,nil
	}
     return false,nil
}

//查询副表
func QueryDetail(url string)([]model.Detail,error){
	var  info model.Info
	var  detail [] model.Detail
    err := model.DB.Table("Infos").Where("url=?",url).First(&info).Error
	if err != nil {
		return detail, err
	}

	id := info.ID
   result := model.DB.Where("index_id=?",id).Find(&detail)
   if result.Error != nil {
	   return detail, result.Error
   }
    return detail, nil
}
