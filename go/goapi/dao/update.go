package dao

import (
	
	"project/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 保存数据
func UpdateInfo(info model.Info) error {
	return model.DB.Transaction(func(tx *gorm.DB) error {
		var deal model.Info
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("url=?", info.Url).First(&deal).Error; err != nil {
			return err
		}
       /* result:=tx.Debug().Model(&deal).Updates(info)
		fmt.Println("影响行数",result.RowsAffected)
		*/
		if err := tx.Model(&deal).Updates(info).Error; err != nil {
			return err
		}
		if len(info.Details) > 0 {
			for i := range info.Details {
				info.Details[i].IndexId = deal.ID
			}
		
		if err := tx.Create(&info.Details).Error; err != nil {
			return err
		}
	}
		return nil
	})	
}
//更新频率和下次爬取时间
func UpdateDoNext(url string ,do int, next int64)error{
   update :=map[string]interface{}{
	 "DoMinute":do,
	 "NextCrawlerTime":next,
   }
 err:= model.DB.Model(&model.Info{}).Where("url=?",url).Updates(update).Error
  if err!=nil{
	return err
  }
  return nil
}