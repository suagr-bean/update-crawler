package dao

import (
	"fmt"
	"project/model"
	"project/pkg"

	"testing"
)

func TestQuery(t *testing.T) {
	path := "../data/test.db"
	db, err := pkg.DBInit(path)
	model.DB = db
	if err != nil {
		fmt.Println(err)
		return
	}
	show := model.Show{}
	show.Url = "https://feeds.simplecast.com/dHoohVNH"
	show.Size = 2

	result, err := QueryDetail(show)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range result.Details {
		fmt.Printf("标题%v\n链接%v\n", v.Title, v.Link)
	}
	fmt.Println("count", result.Count)
}
