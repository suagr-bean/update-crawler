package detect

import (
	"encoding/xml"
	"io"
)

type Result struct {
	Url        string
	Lastupdate string
	Version    string
	Name       string
}
type XmlVersion struct {
	XMLName xml.Name
	Version string `xml:"version,attr"`
}

// xml检测 看看是什么版本的 rss还是feed
func Detect(url string) (Result, error) {
	var result Result
	 context,_:=CrawlerReq(url)
	 body,_:=io.ReadAll(context.Body)
	version := &XmlVersion{}
	err:= xml.Unmarshal(body, version)
	if err != nil {
		return result, err
	}

	v := version.Do()
	switch v {
	case "RSS2.0":
		rss := &Rss{Version: "RSS2.0"}
		rss.AccessInfo(context)
		u := rss.UpdateTitle()
		result = Result{
			Url:        url,
			Lastupdate: u,
			Version:    "RSS2.0"}
	}
	return result, nil
}

// 判断是rss还是feed
func (v *XmlVersion) Do() string {
	switch v.XMLName.Local {
	case "rss":
		if v.Version == "2.0" {
			return "RSS2.0"
		} else {
			return "RSS_Old"
		}
	case "feed":
		return "ATOM"
	case "RDF":
		return "RSS1"
	default:
		return "nil"
	}
}
