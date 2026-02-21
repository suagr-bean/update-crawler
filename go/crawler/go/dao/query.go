package dao

import "project/model"
//查询主数据库
func QueryUrl(show model.Show)([]model.Info,error){
	var info [] model.Info
	if show.Size <= 0 {
		show.Size = 30
	}
	if show.Start <= 0 {
		show.Start = 1
	}
	offset := (show.Start-1) * show.Size
	err := model.DB.Limit(show.Size).Offset(offset).Find(&info).Error
    if err != nil {
		return info, err
	}
	return info, nil
}
//查询主数据库有无数据
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
func QueryDetail(show model.Show)([]model.Detail,error){
	var  info model.Info
	var  detail [] model.Detail
    err := model.DB.Table("Infos").Where("url=?",show.Url).First(&info).Error
	if err != nil {
		return detail, err
	}
    if show.Size<=0{
		show.Size=50
	}
	if show.Start<0{
		show.Start=0
	}
	offset:=show.Size*show.Start
	id := info.ID
   result := model.DB.Limit(show.Size).Offset(offset).Where("index_id=?",id).Order("published_time Desc").Find(&detail)
   if result.Error != nil {
	   return detail, result.Error
   }
    return detail, nil
}
//查询副表文章模糊查询
func QuerySeach(show model.Show)([]model.Detail,error){
	var info model.Info
   var detail [] model.Detail
	if show.Size<=0{
		show.Size=15
	}
	if show.Start<=0{
		show.Start=1
	}
	offset:= (show.Start-1)*show.Size
	err:=model.DB.Where("url=?",show.Url).First(&info).Error
	if err!=nil{
		return detail,err 
	}
	err=model.DB.Where("index_id=?AND title Like?",info.ID,"%"+show.SeachTitle+"%").Order("published_time Desc").Limit(show.Size).Offset(offset).Find(&detail).Error
	if err!=nil{
		return detail,err
	}
    return detail,nil
}
