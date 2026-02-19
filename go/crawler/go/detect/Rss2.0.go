package detect

import (
	
	"encoding/xml"
	"project/model"
	"fmt"
)

type RssProcess struct {
	Data *model.Rss
}

func(r*RssProcess)Deal(data[]byte)model.DealData{
	if len(data)==0{
      return model.DealData{}
	}
	err:=xml.Unmarshal(data,&r.Data) 
    if err!=nil{
		fmt.Println(err)
		return model.DealData{}
	}
	all:=[]model.Article{}
	for _,v:=range r.Data.Items{
	  time:= DealTime(v.PublishTime)
	 article:= model.Article{
        PublishTime:time,
        Title:v.Title,
		AutoLink:v.Enclosure.URL,
		Link:v.Link,
		Guid:v.Guid,
	  }
      all=append(all,article)
	}
     deal:=model.DealData{
		Name:r.Data.GetName(),
		Version:"RSS",
		LastUpdate:r.Data.Last(),
        Articles:all,
	 }
	 return deal
}