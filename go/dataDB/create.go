package dataDB

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 创建主表数据
func (d *DBData) Create() error {
	result := db.Create(&d)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("影响", result.RowsAffected)
	return nil
}

// 副表数据 item的文章啥的 g 就是init里的全局db
func (d *Detail) SaveInfo(g *gorm.DB, url string) error {
	//事务
	return g.Transaction(func(tx *gorm.DB) error {
		var parent DBData
		if err := tx.Select("id").Where("url=?", url).First(&parent).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println("回滚")
				return err
			} else {
				return err
			}
		}
		d.DBDataId = parent.ID
		if err := tx.Create(&d).Error; err != nil {
			return err
		}
		return nil
	})

}
