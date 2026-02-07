package dataDB

import "fmt"

//查询数据库有没有name
func Query(name string)(DBData,error){
	var data DBData
	result:=db.Where("name=?",name).First(&data)
	if result.Error!=nil{

	}	
	return data,nil
}
//查询数据库所有数据
func QueryAll()([]DBData,error){
	var data [] DBData
	gorm:=db.Find(&data)
   if gorm.Error!=nil{
     return nil, gorm.Error
   }
   if gorm.RowsAffected==0{

	 return nil,fmt.Errorf("nodata%v",gorm.RowsAffected)
   }
   return  data,nil
}
