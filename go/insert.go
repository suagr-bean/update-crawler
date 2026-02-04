package main

import (
	"database/sql"
	"fmt"
)

func Insert(task []Task, db *sql.DB) {
	query := `INSERT OR IGNORE INTO crawler (name,url,etag,modified,hashcode) VALUES(?,?,?,?,?)`
	for i := 0; i < len(task); i++ {
		name := task[i].Name
		url := task[i].Url
		etag := task[i].Etag
		hashcode := task[i].Hash
		modified := task[i].Last
		d, err := db.Exec(query, name, url, etag, modified, hashcode)
		if err != nil {
			fmt.Printf("错误%v",err)
			continue
		}
		fmt.Println(d)
	}
}
