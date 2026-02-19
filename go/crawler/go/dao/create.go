package dao

import "project/model"
func  Create( info* model.Info)error{
  err:=model.DB.Create(&info).Error
  if err!=nil{
    return err
  }
  return nil
}