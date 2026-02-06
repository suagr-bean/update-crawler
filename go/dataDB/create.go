package dataDB 
import "fmt"
//创建数据 
func (d*DBData)Create ()error{
	result:=db.Create(&d)
	if result.Error!=nil{
		return result.Error
	}
	fmt.Println("影响",result.RowsAffected)
	return nil
}