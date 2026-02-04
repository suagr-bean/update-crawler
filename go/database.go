package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" //驱动
)

// 数据库初始化连接
func DataBase() *sql.DB {
	db, _ := sql.Open("sqlite", "data.db")
	//表不存在的话创建
	query := `
  CREATE TABLE IF NOT EXISTS crawler(
  id INTEGER PRIMARY KEY AUTOINCREMENT ,
  name TEXT,
  url TEXT UNIQUE NOT NULL,
   etag TEXT,
   modified TEXT,
  hashcode TEXT);`
	_, err := db.Exec(query)
	if err != nil {
		log.Println("建表失败")
	}
	return db
}
