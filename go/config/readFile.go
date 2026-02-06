package config

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)
type CsvData struct{
	Name string 
	Url string
}
//读用户输入的配置文件存URL
func ReadFile(path string) ([]CsvData,error){
	 var csvs []CsvData
    file,err:=os.Open(path)
	if err!=nil{
		return csvs, err
	}
	defer file.Close()
	
   read:= csv.NewReader(file)
    _,err=read.Read()
	 if err!=nil{
		return csvs,err
	 }
   for {
	data,err:= read.Read()
	if err==io.EOF{
	  break
	}
    if err!=nil{
		return csvs, err
	}
    csvs=append(csvs,CsvData{
		Name: data[0],
		Url: data[1],
	})
   }
   //check user not set data
   if len(csvs) <=0{
	err=fmt.Errorf("path no data")
	return csvs,err
   }
	return csvs,nil
}