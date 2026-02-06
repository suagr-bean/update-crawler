package dataDB
//查询数据库有没有name
func Query(name string)(bool,error){
	var count int64
   err:=db.Model(&DBData{}).Where("name=?",name).Count(&count).Error
   if err!=nil{
     
	 return  false,err
   }
    if count>0{
		return true,nil
	}else {
		return false,nil
	}
}