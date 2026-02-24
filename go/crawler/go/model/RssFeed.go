package model

type Dealing interface {
	Deal(data []byte) DealData
}
type DealData struct {
	Name       string
	Version    string
	LastUpdate string
	Articles   []Article
}
type Article struct {
	Title       string
	AutoLink    string
	Link        string
	PublishTime int64
	Guid        string
}
type Rss struct {
	Version string
	Title   string `xml:"channel>title"`
	Items   []Item `xml:"channel>item"`
}
type Item struct {
	PublishTime string `xml:"pubDate"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Guid        string `xml:"guid"`
	Content     string `xml:"description"`
	Enclosure   struct {
		URL string `xml:"url,attr"`
	} `xml:"enclosure"`
}

func (rss *Rss) GetName() string {
	return rss.Title
}
func (rss *Rss) Last() string {
	return rss.Items[0].Title
}
