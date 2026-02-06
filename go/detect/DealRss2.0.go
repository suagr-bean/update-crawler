package detect

import ("encoding/xml"
"fmt")

type  Dealing interface {
	 cleaning(data []byte)
}
type Rss struct {
    Title string `xml:"channel>title"`
	Items [] Item `xml:"channel>item"`
}
type Item struct {
	Title string  `xml:"title"`
	Link string `xml:"link"`
	Enclosure struct{
    URL string `xml:"url,attr"`
	} `xml:"enclosure"`
}

func (rss*Rss)cleaning(data []byte){

	 xml.Unmarshal(data,&rss)
	 fmt.Println(rss.Title)
	 fmt.Println(rss.Items[0].Title)
	 fmt.Println(rss.Items[0].Enclosure.URL)
}