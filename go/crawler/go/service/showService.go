package service

import (
	"project/dao"
	"project/model"
	
	"log"
)

func ShowService( show model.Show)[]model.ShowResult{
    result, err := dao.QueryUrl(show)
	if err != nil {
		log.Printf("Error querying URL: %v", err)
		return []model.ShowResult{} // Return empty slice on error
	}

	var r [] model.ShowResult

    for _,v:=range result{
       showresult:=model.ShowResult{
		Name:v.Name,
		LastUpdate:v.LastUpdate,
		Url: v.Url,
	   }
       r=append(r,showresult)
	}
	return r
}