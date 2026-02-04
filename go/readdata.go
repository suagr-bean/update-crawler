package main

import "database/sql"
import "fmt"

type Data struct {
	Id       int
	Name     string
	Url      string
	HashCode string
	Etag     string
	Modified string
}

// 读数据库 开始爬
type Update struct {
	Name string
	Etag string
	Modified string
}

func ReadData(db *sql.DB) {
	var update []Update
	row, _ := db.Query("SELECT *FROM crawler")
	defer row.Close()
	for row.Next() {
		var data Data
		row.Scan(&data.Id, &data.Name, &data.Url, &data.HashCode, &data.Etag, &data.Modified)
		if data.Etag != "" {
			 newEtag:= CrawlerDeal(data) //处理 etag
			if newEtag!="" {
				fmt.Println("更新")
				update = append(update, Update{Name: data.Name, Etag: newEtag})
			} else {
				fmt.Println("没更新")
			}
		} else if data.Modified != "" {
            newMod:=CrawlerDeal(data)
			if newMod!=""{
				fmt.Println("更新")
				update=append(update,Update {Name:data.Name,Modified:newMod})
			}else {
				fmt.Println("没更新")
			}
		}

	}
	for _, v := range update {
		Updatedata(v.Etag,v.Modified, v.Name, db)
	}
}
