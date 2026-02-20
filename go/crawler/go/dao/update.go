package dao

import (
	"project/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
//保存数据
func UpdateInfo(info model.Info)error{
	return model.DB.Transaction(func (tx*gorm.DB)error{
		var deal model.Info
		if err:=tx.Clauses(clause.Locking{Strength:"UPDATE"}).Where("url=?",info.Url).First(&deal).Error;err!=nil{
			return err
		}
		if err:=tx.Model(&deal).Update("last_update",info.LastUpdate).Error;err!=nil{
			return err
		}
		if len(info.Details)>0{
			for i:=range info.Details{
				info.Details[i].IndexId=deal.ID
			}
		}
		if err:=tx.Create(&info.Details).Error;err!=nil{
			return err
		}
		return nil
	})
    return nil
}