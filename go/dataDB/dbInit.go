package dataDB
import (
	"gorm.io/gorm"
	 "gorm.io/driver/sqlite"
)
type DBData struct{
	gorm.Model
	Name string `gorm:"type:string;unique;notnull"`
	Url string   `gorm:"type:text;unique;notnull"`
	Version string  `gorm:"type:string"`
	LastUpdate string `gorm:"type:text"`
}
 var db *gorm.DB
//初始化数据库
func InitDB()error{
	var err error
	 db,err=gorm.Open(sqlite.Open("xml.db"),&gorm.Config{})
	 if err!=nil{
       return err
	 }
	 db.AutoMigrate(DBData{})
	 return nil
}