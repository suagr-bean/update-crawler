package detect

import (
	"encoding/xml"
	
	"project/model"
)

type XmlVersion struct {
	XMLName xml.Name
	Version string `xml:"version,attr"`
}

// 探测
func Detect( cont model.Context) (model.DealData, error) {
	var data model.DealData
	var version string
	resp, err := Crawler(cont)
	if err != nil {
		return data, err
	}
	if resp.Code==304{
		return data,nil
	}
	
    if cont.Version==""{
	v := &XmlVersion{}
	err = xml.Unmarshal(resp.Body, v)
	if err != nil {
		return data, err
	}
	version= v.Do()
} else {
	version=cont.Version
}
     
	switch version{
	case "RSS2.0":
		rss := RssProcess{}
		data,_= rss.Deal(resp.Body)
	
	}
	data.Etag=cont.Etag
	data.LastModified=cont.LastModified
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
