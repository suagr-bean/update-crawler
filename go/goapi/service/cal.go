package service

import (
	"project/dao"
	"project/model"
	"fmt"
	"time"
)

func Caltimes(url string) int{
	//dao层拿到30篇文章
	show:=model.Show{
		Start:1,
		Size:30,
		Url:url,
	}
	
	hourly:=make([]int,24)
	//判断 有无最新30条
	result,_:=dao.QueryDetail(show)
	fmt.Println(len(result.Details))
	if len(result.Details)>=30{
	   for i:=0;i<len(result.Details);i++{
		 t:=time.Unix(result.Details[i].PublishedTime,0)
		 hour:=t.Hour()
         hourly[hour]++
	   
	  }
	}else{
      return 12*60
	}
	max:=hourly[0]
	index:=0//计数桶的时间
	for i,v:=range hourly{
		if v>max{
			index=i
			max=v
		}
	}
	return index*60 //返回分钟
}