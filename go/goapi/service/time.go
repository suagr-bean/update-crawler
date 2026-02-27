package service

import (
	"project/Utils"
	"project/model"
	"time"
)

//爬虫时间算法
func DoTime( cal model.TimeCal)model.TimeCal {
    var  new int 
	var next time.Time
	//更新
    if cal.IsUpdate ==true{
	  new= int(float64(cal.Interval)*0.7)
	  
	}else {
		new =int(float64(cal.Interval)*1.3)

	}
     next=Utils.DealTime(new)
	 return model.TimeCal{
		Next:next,
		Interval: new,
	 }
}