package detect

import (
	"encoding/xml"
	"io"

	"fmt"
	"net/http"
)
type XmlVersion struct {
	XMLName xml.Name 
	Version string `xml:"version,attr"`
}
//xml检测 看看是什么版本的 rss还是feed
func  Detect(url string) (string,error){
    client:=&http.Client{}
	req,err:=http.NewRequest("GET",url,nil)
	if err!=nil{
      return "",err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept","application/xml,text/xml,*/*")
     resp,err:=client.Do(req)
	 if err!=nil{
		return "",err
	 }
	 defer resp.Body.Close()
	body,_:= io.ReadAll(resp.Body)
	
    version:=&XmlVersion{}
   err=xml.Unmarshal(body,version)
   if err!=nil{
      return "",err
   }
     
	  result:=version.Do()
	  fmt.Println(result)
      return result,nil
	 
}
//判断是rss还是feed
func (v*XmlVersion)Do()string{
    switch v.XMLName.Local{
	case "rss":
	if v.Version=="2.0"{
     return "RSS2.0"
	}else{
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