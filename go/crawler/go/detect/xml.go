package detect

import (
	"encoding/xml"
	"io"
	"project/model"
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

// 探测
func Detect(url string) (model.DealData, error) {
	var data model.DealData
	resp, err := Crawler(url)
	if err != nil {
		return data, err
	}
	if resp.Body == nil {
		// return data, errors.New("response body is nil") // More specific error
		return data, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	version := &XmlVersion{}
	err = xml.Unmarshal(body, version)
	if err != nil {
		return data, err
	}

	v := version.Do()
	switch v {
	case "RSS2.0":
		rss := RssProcess{}
		data = rss.Deal(body)
		return data, nil
	}
	return data, nil
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
