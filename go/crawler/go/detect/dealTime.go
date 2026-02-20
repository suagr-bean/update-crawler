package detect

import (
	"strings"
	"time"
)

//处理时间
func DealTime(timeStr string) int64{
	//清洗掉空格什么的
	cleanTime:=strings.TrimSpace(timeStr)
	layouts:=[]string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02",
		time.RubyDate,
		time.ANSIC,
	}
    now:=time.Now().Unix()
	for _,v:=range layouts{
     t,err:=time.Parse(v,cleanTime)
	 if err==nil{
      
	  if t.Unix()>now+6000{
		return now
	  }
		return t.Unix()
	 }
	}
	return time.Now().Unix()
}
