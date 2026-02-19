package model


type Dealing interface {
	Deal(data[]byte)DealData
}
type DealData struct{
	Name string
	Version string
	LastUpdate string
	Articles []Article
}
type Article struct{
	Title string 
	AutoLink string
	Link string 

}
type Rss struct {
	Version string
	Title   string `xml:"channel>title"`
	Items   []Item `xml:"channel>item"`
}
type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Content   string `xml:"description"`
	Enclosure struct {
		URL string `xml:"url,attr"`
	} `xml:"enclosure"`
}
func (rss*Rss) GetName()string{
	return rss.Title
}
func (rss*Rss)Last()string{
	return rss.Items[0].Title
}


