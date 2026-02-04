package main

import (
	"database/sql"
	"fmt"
)

func Updatedata(etag string, name string, mod string,db *sql.DB) {
	if db == nil {
		fmt.Println("db为空")
		return
	}
	query := `UPDATE crawler SET etag=?,modified=? WHERE name=?`
	result, err := db.Exec(query, etag,mod,name)
	if err != nil {
		fmt.Printf("问题 %v", err)
	}
	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("有问题")
	}
	if row == 0 {
		fmt.Println("插入数据库失败")
	}
}
