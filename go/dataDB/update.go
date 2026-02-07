package dataDB
//更新数据 lastupdate
func(d*DBData) UpdateLast(){
	  name:=d.Name
       last:=d.LastUpdate
	db.Model(&DBData{}).Where("name=?",name).Update("LastUpdate",last)
}