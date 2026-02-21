package detect

import (
	"strings"
	"time"
	"fmt"
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
		"Mon, 2 Jan 2006 15:04:05 -0700",
		"Mon, 2 Jan 2006 15:04:05 MST",
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
	
	fmt.Printf("错误时间%",cleanTime)
	return 0
	//return time.Now().Unix()
}
