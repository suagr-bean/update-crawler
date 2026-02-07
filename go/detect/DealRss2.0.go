package detect

import (
	"encoding/xml"
	"io"
	"fmt"
)

type Dealing interface {
	AccessInfo(context Context)
	UpdateTitle() string
}
type Rss struct {
	Version string
	Title   string `xml:"channel>title"`
	Items   []Item `xml:"channel>item"`
}
type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Enclosure struct {
		URL string `xml:"url,attr"`
	} `xml:"enclosure"`
}

func (rss *Rss) AccessInfo(context Context) {
	 defer context.Body.Close()
     data,err:=io.ReadAll(context.Body)
	 if err!=nil{
		fmt.Println("wrong")
		return 
	 }
	 err=xml.Unmarshal(data, rss)
     if err!=nil{
		fmt.Println("解析错误")
	 }
}
func (rss *Rss) UpdateTitle() string {
	return rss.Items[0].Title
}
