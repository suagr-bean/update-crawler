package main

import "database/sql"
import _"modernc.org/sqlite"
func ReadSQL()(*sql.DB,error){
	data,_:=sql.Open("sqlite",".crawler.db")
	
	creatTable:=`CREATE TABLE IF NOT EXISTS crawler( 
	url  TEXT PRIMARY KEY, strate INTEGER,created_at DATEIME DEFAULT CURRENT_TIMESTAMP);`
	 _,err:=data.Exec(creatTable);
	 if err!=nil{
		return nil, err
	 }
    return data,nil
}